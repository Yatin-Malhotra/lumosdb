package server

import (
	"log"
	"net"
)

type Server struct {
	Addr string
}

func New(addr string) *Server {
	return &Server{Addr: addr}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	log.Printf("Server is listening on %s", s.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		go s.handleConnection(conn)
	}
}
