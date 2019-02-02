package main

import (
	"docsdocs/log"
	"docsdocs/proto"
)

func main() {
	log.Settings("json", "stdout", "debug")
	logger := log.NewDocsLogger()
	logger.Info("Starting client")
	t, err := proto.NewTransport("tcp", "localhost:5000")
	if err != nil {
		logger.Error(err)
	}
	resp, err := t.RoundTrip(&proto.Request{
		Header:    proto.Header{},
		Method:    proto.MethodGet,
		BodyBytes: []byte("hello =]"),
	})
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Response: %v\n\n", resp)
}
