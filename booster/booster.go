package booster

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/danielmorandini/booster/network"
	"github.com/danielmorandini/booster/network/packet"
	"github.com/danielmorandini/booster/node"
	"github.com/danielmorandini/booster/protocol"
	"github.com/danielmorandini/booster/socks5"
)

type Proxy interface {
	NotifyTunnel() (<-chan interface{}, error)
	ListenAndServe(ctx context.Context, port int) error
}

type Conn struct {
	network.Conn

	ID         string
	RemoteNode *node.Node
}

type Network struct {
	LocalNode *node.Node
	Conns     []*Conn
}

// Booster wraps the parts that compose a booster node together.
type Booster struct {
	*log.Logger

	Proxy Proxy

	mux     sync.Mutex
	Network *Network

	netconfig network.Config
	stop      chan struct{}
}

// New creates a new configured booster node. Creates a network configuration
// based in the information contained in the protocol package.
//
// The internal proxy is configured to use the node dispatcher as network
// dialer.
func New(pport, bport int) (*Booster, error) {
	b := new(Booster)

	log := log.New(os.Stdout, "BOOSTER  ", log.LstdFlags)
	dialer := node.NewDispatcher(b.Nodes)
	proxy := socks5.New(dialer)
	netconfig := network.Config{
		TagSet: packet.TagSet{
			PacketOpeningTag:  protocol.PacketOpeningTag,
			PacketClosingTag:  protocol.PacketClosingTag,
			PayloadClosingTag: protocol.PayloadClosingTag,
			Separator:         protocol.Separator,
		},
	}
	pp := strconv.Itoa(pport)
	bp := strconv.Itoa(bport)
	node, err := node.New("localhost", pp, bp, true)
	if err != nil {
		return nil, err
	}
	network := &Network{
		LocalNode: node,
		Conns:     []*Conn{},
	}

	b.Logger = log
	b.Proxy = proxy
	b.Network = network
	b.netconfig = netconfig
	b.stop = make(chan struct{})

	return b, nil
}

// Run starts the proxy and booster node.
//
// This is a blocking routine that can be stopped using the Close() method.
// Traps INTERRUPT signals.
func (b *Booster) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errc := make(chan error, 2)
	_, pport, _ := net.SplitHostPort(b.Network.LocalNode.PAddr.String())
	_, bport, _ := net.SplitHostPort(b.Network.LocalNode.BAddr.String())
	pp, _ := strconv.Atoi(pport)
	bp, _ := strconv.Atoi(bport)
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		errc <- b.ListenAndServe(ctx, bp)
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		errc <- b.Proxy.ListenAndServe(ctx, pp)
		wg.Done()
	}()

	// trap exit signals
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		for sig := range c {
			b.Printf("booster: signal (%v) received: exiting...", sig)
			b.Close()
			return
		}
	}()

	select {
	case err := <-errc:
		cancel()
		wg.Wait()
		return err
	case <-b.stop:
		cancel()
		wg.Wait()
		return fmt.Errorf("booster: stopped")
	}
}

// Close stops the Run routine. It drops the whole booster network, preparing for the
// node to reset or stop.
func (b *Booster) Close() error {
	b.stop <- struct{}{}
	return nil
}

// ListenAndServe shows to the network, listening for incoming tcp connections an
// turning them into booster connections.
func (b *Booster) ListenAndServe(ctx context.Context, port int) error {
	p := strconv.Itoa(port)
	ln, err := network.Listen("tcp", ":"+p, b.netconfig)
	if err != nil {
		return err
	}
	defer ln.Close()

	b.Printf("listening on port: %v", p)

	errc := make(chan error)
	defer close(errc)

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				errc <- fmt.Errorf("booster: cannot accept conn: %v", err)
				return
			}

			go b.Handle(ctx, conn)
		}
	}()

	select {
	case err := <-errc:
		return err
	case <-ctx.Done():
		ln.Close()
		<-errc // wait for listener to return
		return ctx.Err()
	}
}

func (b *Booster) Handle(ctx context.Context, conn *network.Conn) {
	defer conn.Close()

	pkts, err := conn.Consume()
	if err != nil {
		b.Printf("booster: cannot consume packets: %v", err)
		return
	}

	b.Println("booster: consuming packets...")

	for p := range pkts {
		b.Printf("booster: consuming packet: %+v", p)
	}

	b.Println("booster: packets consumed.")
}

func (b *Booster) Nodes() (*node.Node, []*node.Node) {
	b.mux.Lock()
	defer b.mux.Unlock()

	root := b.Network.LocalNode
	nodes := []*node.Node{}

	for _, c := range b.Network.Conns {
		nodes = append(nodes, c.RemoteNode)
	}

	return root, nodes
}

func (b *Booster) DialContext(ctx context.Context, netwk, addr string) (*network.Conn, error) {
	// first configure a dialer and create a new connection
	dialer := network.NewDialer(new(net.Dialer), b.netconfig)
	conn, err := dialer.DialContext(ctx, netwk, addr)
	if err != nil {
		return nil, err
	}

	return conn, err
}
