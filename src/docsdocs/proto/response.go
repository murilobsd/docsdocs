package proto

import (
	"bytes"
	"encoding/gob"
)

// Status ...
type Status int64

const (
	// Sucess ...
	Sucess Status = 1 << iota
	// Error ...
	Error
)

// Response ...
type Response struct {
	Header    Header
	Status    Status
	BodyBytes []byte
}

// Serialize - convert the response to the wire format
func (r *Response) Serialize() ([]byte, error) {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(r)
	return b.Bytes(), err
}

// DeserializeResponse - convert the response from the wire format
func DeserializeResponse(responseBytes []byte) (Response, error) {
	b := bytes.Buffer{}
	b.Write(responseBytes)
	decoder := gob.NewDecoder(&b)
	resp := &Response{}
	err := decoder.Decode(resp)
	return *resp, err
}
