package node

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"github.com/danielmorandini/booster/network"
	"github.com/danielmorandini/booster/socks5"
)

type Proxy interface {
	NotifyTunnel() (<-chan interface{}, error)
	ListenAndServe(ctx context.Context, port int) error
}

type Booster struct {
	*log.Logger

	Proxy Proxy

	stop chan struct{}
}

func NewBooster() *Booster {
	log := log.New(os.Stdout, "BOOSTER  ", log.LstdFlags)
	dialer := new(net.Dialer)
	proxy := socks5.New(dialer)

	b := new(Booster)
	b.Logger = log
	b.Proxy = proxy
	b.stop = make(chan struct{})

	return b
}

func (b *Booster) Run(pport, bport int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errc := make(chan error)
	go func() {
		errc <- b.ListenAndServe(ctx, bport)
	}()

	go func() {
		errc <- b.Proxy.ListenAndServe(ctx, pport)
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
		<-errc // wait for the other goroutine to return
		return err
	case <-b.stop:
		cancel()
		<-errc // wait for ListenAndServe to return
		<-errc // wait for ListenAndServe to return
		return fmt.Errorf("booster: stopped")
	}
}

func (b *Booster) Close() error {
	b.stop <- struct{}{}
	return nil
}

func (b *Booster) ListenAndServe(ctx context.Context, port int) error {
	p := strconv.Itoa(port)
	ln, err := network.Listen("tcp", ":"+p)
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
