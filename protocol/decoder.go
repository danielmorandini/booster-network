package protocol

import (
	"fmt"
	"reflect"

	"github.com/danielmorandini/booster/protocol/internal"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

// DecoderFunc defines how a decoder should behave.
type DecoderFunc func([]byte) (interface{}, error)

// Implemented default decoders
var decoders = map[Message]DecoderFunc{
	MessageHello:      decodeHello,
	MessageCtrl:       decodeCtrl,
	MessageBandwidth:  decodeBandwidth,
	MessageInspect:    decodeInspect,
	MessageConnect:    decodeConnect,
	MessageDisconnect: decodeDisconnect,
	MessageNode:       decodeNode,
	MessageHeartbeat:  decodeHeartbeat,
	MessageTunnel:     decodeTunnelEvent,
}

// Decoder wraps a list of implemented decoder functions.
type Decoder struct {
	Decoders map[Message]DecoderFunc
}

// NewDecoder returns new instance of Decoder, filled with a default list
// of decoder functions ready to be used.
func NewDecoder() *Decoder {
	return &Decoder{
		Decoders: decoders,
	}
}

// Decode takes as input a byte slice and tries to decode it into v.
// msg is used to determine which decoding function should be used
// internally.
//
// v has to be a pointer to a struct.
func (d Decoder) Decode(p []byte, v interface{}, msg Message) error {
	// first find the right decoder function for this message (if any)
	f, ok := d.Decoders[msg]
	if !ok {
		return fmt.Errorf("protocol: decode error: could find any decode function for message (%v)", msg)
	}

	// perform the decoding
	s, err := f(p)
	if err != nil {
		return fmt.Errorf("protocol: decode error: %v", err)
	}

	// reflect the actual value decoded
	val := reflect.ValueOf(s)

	// reflect the actual value of the interface provided. Need to retrieve its pointer
	// in order to be able to set to it
	ptr := reflect.ValueOf(v).Elem()
	if !ptr.CanSet() {
		return fmt.Errorf("protocol: unable to set on value, pass pointer to struct instead")
	}

	// Copy contents of the decoded payload into the parameter
	// check if we're talking about the same thing
	if ptr.Type() != val.Type() {
		return fmt.Errorf("protocol: decode error: trying to reflect %v into %v, which is illegal", ptr.Type(), val.Type())
	}

	ptr.Set(val)
	return nil
}

// TODO: probably this whole boilerplate code below could be replaced
// using reflection.

func decodeHello(p []byte) (interface{}, error) {
	payload := new(internal.PayloadHello)
	if err := proto.Unmarshal(p, payload); err != nil {
		return nil, err
	}

	return &PayloadHello{
		BPort: payload.Bport,
		PPort: payload.Pport,
	}, nil
}

func decodeCtrl(p []byte) (interface{}, error) {
	payload := new(internal.PayloadCtrl)
	if err := proto.Unmarshal(p, payload); err != nil {
		return nil, err
	}

	return &PayloadCtrl{
		Operation: Operation(payload.Operation),
	}, nil
}

func decodeBandwidth(p []byte) (interface{}, error) {
	payload := new(internal.PayloadBandwidth)
	if err := proto.Unmarshal(p, payload); err != nil {
		return nil, err
	}

	return &PayloadBandwidth{
		Tot:       int(payload.Tot),
		Bandwidth: int(payload.Bandwidth),
		Type:      payload.Type,
	}, nil
}

func decodeInspect(p []byte) (interface{}, error) {
	payload := new(internal.PayloadInspect)
	if err := proto.Unmarshal(p, payload); err != nil {
		return nil, err
	}

	features := []Message{}
	for _, v := range payload.Features {
		features = append(features, Message(v))
	}

	return &PayloadInspect{
		Features: features,
	}, nil
}

func decodeConnect(p []byte) (interface{}, error) {
	payload := new(internal.PayloadConnect)
	if err := proto.Unmarshal(p, payload); err != nil {
		return nil, err
	}

	return &PayloadConnect{
		Target: payload.Target,
	}, nil
}

func decodeDisconnect(p []byte) (interface{}, error) {
	payload := new(internal.PayloadDisconnect)
	if err := proto.Unmarshal(p, payload); err != nil {
		return nil, err
	}

	return &PayloadDisconnect{
		ID: payload.Id,
	}, nil
}

func decodeNode(p []byte) (interface{}, error) {
	payload := new(internal.PayloadNode)
	if err := proto.Unmarshal(p, payload); err != nil {
		return nil, err
	}

	ts := []*Tunnel{}
	for _, t := range payload.Tunnels {
		tunnel := &Tunnel{
			ID:     t.Id,
			Target: t.Target,
			Acks:   int(t.Acks),
			Copies: int(t.Copies),
		}

		ts = append(ts, tunnel)
	}

	return &PayloadNode{
		ID:      payload.Id,
		BAddr:   payload.Baddr,
		PAddr:   payload.Paddr,
		Active:  payload.Active,
		Tunnels: ts,
	}, nil
}

func decodeHeartbeat(p []byte) (interface{}, error) {
	payload := new(internal.PayloadHeartbeat)
	if err := proto.Unmarshal(p, payload); err != nil {
		return nil, err
	}

	t, err := ptypes.Timestamp(payload.Ttl)
	if err != nil {
		return nil, err
	}

	return &PayloadHeartbeat{
		ID:   payload.Id,
		Hops: int(payload.Hops),
		TTL:  t,
	}, nil
}

func decodeTunnelEvent(p []byte) (interface{}, error) {
	payload := new(internal.PayloadTunnelEvent)
	if err := proto.Unmarshal(p, payload); err != nil {
		return nil, err
	}

	return &PayloadTunnelEvent{
		Target: payload.Target,
		Event:  int(payload.Event),
	}, nil
}