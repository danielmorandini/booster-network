package booster

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/danielmorandini/booster/network"
	"github.com/danielmorandini/booster/network/packet"
	"github.com/danielmorandini/booster/node"
	"github.com/danielmorandini/booster/protocol"
	"github.com/danielmorandini/booster/pubsub"
)

type Networks map[string]*Network

var Nets = &Networks{}

func (n Networks) Get(id string) *Network {
	net, ok := n[id]
	if !ok {
		panic("networks: tried to get unregistered network: " + id)
	}

	return net
}

func (n Networks) Set(id string, net *Network) {
	_, ok := n[id]
	if ok {
		panic("networks: tried to set already registered network: " + id)
	}

	net.boosterID = id
	n[id] = net
}

// Network describes a booster network: a local node, connected to other booster nodes
// using network.Conn as connector.
type Network struct {
	*log.Logger
	PubSub

	boosterID string

	mux       sync.Mutex
	LocalNode *node.Node
	Conns     map[string]*Conn
}

func NewNet(n *node.Node, boosterID string) *Network {
	return &Network{
		Logger:    log.New(os.Stdout, "NETWORK  ", log.LstdFlags),
		PubSub:    pubsub.New(),
		LocalNode: n,
		boosterID: boosterID,
		Conns:     make(map[string]*Conn),
	}
}

func (n *Network) AddConn(c *Conn) error {
	n.mux.Lock()
	defer n.mux.Unlock()

	if _, ok := n.Conns[c.ID]; ok {
		return fmt.Errorf("network: conn (%v) already present", c.ID)
	}

	c.boosterID = n.boosterID
	n.Conns[c.ID] = c
	return nil
}

func (n *Network) Notify() (chan interface{}, error) {
	return n.Sub(TopicNodes)
}

func (n *Network) StopNotifying(c chan interface{}) {
	n.Unsub(c, TopicNodes)
}

func (n *Network) Nodes() (*node.Node, []*node.Node) {
	n.mux.Lock()
	defer n.mux.Unlock()

	root := n.LocalNode
	nodes := []*node.Node{}

	for _, c := range n.Conns {
		nodes = append(nodes, c.RemoteNode)
	}

	return root, nodes
}

func (n *Network) Ack(node *node.Node, id string) error {
	n.Printf("network: acknoledging (%v) on node (%v)", id, node.ID())

	if err := node.Ack(id); err != nil {
		return err
	}

	n.Pub(node, TopicNodes)
	return nil
}

func (n *Network) RemoveTunnel(node *node.Node, id string, acknoledged bool) error {
	n.Printf("booster: removing (%v) on node (%v)", id, node.ID())

	if err := node.RemoveTunnel(id, acknoledged); err != nil {
		return err
	}

	n.Pub(node, TopicNodes)
	return nil
}

func (n *Network) AddTunnel(node *node.Node, target string) {
	n.Printf("booster: adding tunnel (%v) to node (%v)", target, node.ID())

	node.AddTunnel(target)
	n.Pub(node, TopicNodes)
}

func (n *Network) NewConn(conn *network.Conn, node *node.Node, id string) *Conn {
	return &Conn {
		Conn: conn,
		RemoteNode: node,
		ID: id,
		boosterID: n.boosterID,
	}
}

// Conn adds an identifier and a convenient RemoteNode field to a bare network.Conn.
type Conn struct {
	*network.Conn

	ID         string // ID is usually the remoteNode identifier.
	boosterID  string
	RemoteNode *node.Node
}

// Close closes the connection and sets the status of the remote node
// to inactive and removes the connection from the network.
func (c *Conn) Close() error {
	if c.Conn == nil {
		return fmt.Errorf("network: connection is closed")
	}

	if err := c.Conn.Close(); err != nil {
		return err
	}
	c.RemoteNode.SetIsActive(false)

	// Remove the connection only if it is actually part of this network
	if _, ok := Nets.Get(c.boosterID).Conns[c.ID]; ok {
		Nets.Get(c.boosterID).Conns[c.ID].Conn = nil
	}

	return nil
}

func (c *Conn) Send(p *packet.Packet) error {
	if c.Conn == nil {
		return fmt.Errorf("network: connection is closed")
	}
	return c.Conn.Send(p)
}

func (c *Conn) Consume() (<-chan *packet.Packet, error) {
	if c.Conn == nil {
		return nil, fmt.Errorf("network: connection is closed")
	}
	return c.Conn.Consume()
}

func (c *Conn) Recv() (*packet.Packet, error) {
	if c.Conn == nil {
		return nil, fmt.Errorf("network: connection is closed")
	}
	return c.Conn.Recv()
}

func ValidatePacket(p *packet.Packet) error {
	// Find header
	hraw, err := p.Module(protocol.ModuleHeader)
	if err != nil {
		return err
	}

	h, err := protocol.DecodeHeader(hraw.Payload())
	if err != nil {
		return err
	}

	// Check packet version
	if !protocol.IsVersionSupported(h.ProtocolVersion) {
		return fmt.Errorf("packet validation: version (%v) is not supported", h.ProtocolVersion)
	}

	// Check that the information contained in the header reflect the
	// actual content of the packet
	for _, mid := range h.Modules {
		if _, err := p.Module(mid); err != nil {
			return fmt.Errorf("packet validation: %v", err)
		}
	}

	return nil
}

func ExtractHeader(p *packet.Packet) (*protocol.Header, error) {
	if err := ValidatePacket(p); err != nil {
		return nil, fmt.Errorf("booster: discarding invalid packet: %v", err)
	}

	// extract header from packet
	hraw, err := p.Module(protocol.ModuleHeader)
	if err != nil {
		return nil, fmt.Errorf("booster: failed reading module header: %v", err)
	}
	h, err := protocol.DecodeHeader(hraw.Payload())
	if err != nil {
		return nil, fmt.Errorf("booster: failed decoding header: %v", err)
	}

	return h, nil
}

func (b *Booster) composeHeartbeat(pl *protocol.PayloadHeartbeat) (*packet.Packet, error) {
	if pl == nil {
		pl = &protocol.PayloadHeartbeat{
			Hops: 0,
			ID:   "heartbeat", // TODO(daniel): unused field
		}
	}

	pl.Hops++
	pl.TTL = time.Now().Add(b.HeartbeatTTL)

	h, err := protocol.HeartbeatHeader()
	if err != nil {
		return nil, err
	}
	hpl, err := protocol.EncodePayloadHeartbeat(pl)
	if err != nil {
		return nil, err
	}

	// compose the packet
	p := packet.New()
	enc := protocol.EncodingProtobuf
	if _, err := p.AddModule(protocol.ModuleHeader, h, enc); err != nil {
		return nil, err
	}
	if _, err := p.AddModule(protocol.ModulePayload, hpl, enc); err != nil {
		return nil, err
	}

	return p, nil
}

func composeNode(n *node.Node) (*packet.Packet, error) {
	h, err := protocol.NodeHeader()
	if err != nil {
		return nil, err
	}

	tunnels := make([]*protocol.Tunnel, len(n.Tunnels()))
	for _, t := range n.Tunnels() {
		tunnel := &protocol.Tunnel{
			ID:     t.ID(),
			Target: t.Target,
			Acks:   t.Acks(),
			Copies: t.Copies(),
		}

		tunnels = append(tunnels, tunnel)
	}
	param := &protocol.PayloadNode{
		ID:      n.ID(),
		BAddr:   n.BAddr.String(),
		PAddr:   n.PAddr.String(),
		Active:  n.IsActive(),
		Tunnels: tunnels,
	}

	npl, err := protocol.EncodePayloadNode(param)
	if err != nil {
		return nil, err
	}

	p := packet.New()
	enc := protocol.EncodingProtobuf
	if _, err = p.AddModule(protocol.ModuleHeader, h, enc); err != nil {
		return nil, err
	}
	if _, err = p.AddModule(protocol.ModulePayload, npl, enc); err != nil {
		return nil, err
	}

	return p, nil
}
