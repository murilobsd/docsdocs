package server

import (
	"docsdocs/log"
	"docsdocs/proto"
	"encoding/gob"
	"net"
)

type Server struct {
	log.Logger
	listener net.Listener
}

func NewServer(proto, address string) (*Server, error) {
	listener, err := net.Listen(proto, address)
	if err != nil {
		return nil, err
	}
	return &Server{
		Logger:   log.NewDocsLogger(),
		listener: listener,
	}, nil
}

func (s *Server) Run() chan struct{} {
	q := make(chan struct{})
	go func(quit chan struct{}) {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				s.Error(err)
				return
			}
			go s.handlerConn(conn)

		}
	}(q)
	return q
}

func (s *Server) handlerConn(conn net.Conn) {
	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)
	for {
		var req proto.Request
		err := decoder.Decode(&req)
		if err != nil {
			s.Error(err)
			return
		}
		s.Info("Got request: %v\n", req)
		var resp = new(proto.Response)
		switch req.Method {
		case proto.MethodGet:
			s.Info("request is a get method.")
			resp.Status = proto.Sucess
		case proto.MethodPost:
			s.Info("request is a post method.")
			resp.Status = proto.Sucess
		case proto.MethodDelete:
			s.Info("request is a delete method.")
			resp.Status = proto.Sucess
		default:
			s.Warn("unknow request method")
			resp.Status = proto.Error
		}
		encoder.Encode(*resp)
	}
}
