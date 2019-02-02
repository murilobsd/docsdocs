package proto

import (
	"encoding/gob"
	"net"
)

// Header ...
type Header struct{}

type Encoder interface {
	Encode(interface{}) error
}

type Decoder interface {
	Decode(interface{}) error
}

// Transport ...
type Transport struct {
	conn net.Conn
	enc  Encoder
	dec  Decoder
}

// NewTransport - create a new transport structure
func NewTransport(proto, remote string) (*Transport, error) {
	conn, err := net.Dial(proto, remote)
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)
	return &Transport{
		conn: conn,
		enc:  enc,
		dec:  dec,
	}, err
}

// RoundTrip - Implementation of a round tripper interface,
// effectively this is how the request will be serialized,
// and put on the wire, and how the response will be deserialized
func (t *Transport) RoundTrip(request *Request) (Response, error) {
	var response Response
	// serialize request
	if err := t.enc.Encode(request); err != nil {
		return response, err
	}
	// unserialize response
	if err := t.dec.Decode(&response); err != nil {
		return response, err
	}
	return response, nil
}
