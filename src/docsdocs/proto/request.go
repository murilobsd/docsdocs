package proto

import (
	"bytes"
	"context"
	"encoding/gob"
)

func init() {
	gob.Register(Request{})
}

// Request ...
type Request struct {
	Header     Header
	Method     Method
	BodyBytes  []byte
	RemoteAddr string
	ctx        context.Context
}

// Context returns the request context
func (r *Request) Context() context.Context {
	if r.ctx != nil {
		return r.ctx
	}
	return context.Background()
}

// Serialize
func (r *Request) Serialize() ([]byte, error) {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(r)
	return b.Bytes(), err
}

func Deserialize(reqBytes []byte) (Request, error) {
	b := bytes.Buffer{}
	b.Write(reqBytes)
	decoder := gob.NewDecoder(&b)
	req := &Request{}
	err := decoder.Decode(req)
	return *req, err
}
