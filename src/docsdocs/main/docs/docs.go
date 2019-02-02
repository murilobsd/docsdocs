package main

import (
	"docsdocs/log"
	"docsdocs/proto"
	"fmt"
)

func main() {
	log.Settings("json", "stdout", "debug")
	t, err := proto.NewTransport("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	for _, method := range []proto.Method{
		proto.MethodGet,
		proto.MethodPost,
		proto.MethodDelete,
		42,
	} {
		fmt.Println("starting request: ", method)
		resp, err := t.RoundTrip(&proto.Request{
			Header:    proto.Header{},
			Method:    method,
			BodyBytes: []byte("hello =]"),
		})
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
		}
		fmt.Printf("Response: %v\n\n", resp)
	}

}
