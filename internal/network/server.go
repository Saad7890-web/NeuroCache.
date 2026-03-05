package network

import "net"

type Server struct {
	address string
}

func NewServer(addr string) *Server {
	return &Server{
		address: addr,
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
		go handleConnection(conn)
	}
}