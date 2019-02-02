package proto

import (
	"bytes"
	"encoding/gob"
)

func init() {
	gob.Register(Request{})
}

// Request ...
type Request struct {
	Header    Header
	Method    Method
	BodyBytes []byte
}

// Serialize ...
func (r *Request) Serialize() ([]byte, error) {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(r)
	return b.Bytes(), err
}

// DeserualizeRequest ...
func DeserualizeRequest(reqBytes []byte) (Request, error) {
	b := bytes.Buffer{}
	b.Write(reqBytes)
	decoder := gob.NewDecoder(&b)
	req := &Request{}
	err := decoder.Decode(req)
	return *req, err
}
