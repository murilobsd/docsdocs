package server

import (
	"docsdocs/log"
	"docsdocs/proto"
	"encoding/gob"
	"io"
	"net"
)

func NewServer() {
	log.Settings("json", "stdout", "debug")
	logger := log.NewDocsLogger()
	listener, _ := net.Listen("tcp", ":5000")
	for {
		conn, _ := listener.Accept()
		logger.Info("Remote Ip: ", conn.RemoteAddr().String())
		decoder := gob.NewDecoder(conn)
		var req proto.Request
		err := decoder.Decode(&req)
		if err != nil {
			logger.Error(err)
			if err == io.EOF {
				continue
			}
		} else {
			logger.Info("Got request: %v\n", req)
			encoder := gob.NewEncoder(conn)
			encoder.Encode(proto.Response{
				BodyBytes: []byte("i'm here"),
			})
		}
	}
}
