package network

import (
	"net"

	"github.com/Saad7890-web/neurocache/internal/engine"
)

type Server struct {
	address string
	engine *engine.Engine
}

func NewServer(addr string) *Server {
	return &Server{
		address: addr,
		engine: engine.NewEngine(),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn, s.engine)
	}
}