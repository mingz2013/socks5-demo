package main

import (
	"log"
	"net"
)

type Server struct {
}

func (s *Server) Listen(address string) {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go s.handleConn(conn)
	}

}

func (s *Server) handleConn(conn net.Conn) {
	// verify

}
