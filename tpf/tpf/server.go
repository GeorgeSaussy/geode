package tpf

import (
	"errors"
	"io"
	"log"
	"net"
	"strconv"
)

// Server instances contain which ports need forwading.
type Server struct {
	lp int
	pp int
	ln net.Listener
}

// NewServer will create a new Server instance or
// return an error if the ports provided are invalid.
func NewServer(lp int, pp int) (*Server, error) {
	if lp <= 0 {
		return nil, errors.New("the listening port must be a positive number")
	}
	if pp <= 0 {
		return nil, errors.New("the proxied port must be a positive number")
	}
	return &Server{
		lp: lp,
		pp: pp,
		ln: nil,
	}, nil
}

// Serve will start the server.
func (s *Server) Serve() error {
	ln, err := net.Listen("tcp", "localhost:"+strconv.Itoa(s.lp))
	if err != nil {
		return err
	}
	s.ln = ln
	log.Printf("server listening on port %d\n", s.lp)
	for {
		cl, err := s.ln.Accept()
		if err != nil {
			// We do not log this error because the loop will not wait
			// for a new connection. In practice this means there will
			// be zillions of failed client connections per second.
			continue
		}
		px, err := net.Dial("tcp", "localhost:"+strconv.Itoa(s.pp))
		if err != nil {
			log.Printf("could not connect to backend: %s\n", err.Error())
			cl.Close()
			continue // return
		}
		defer cl.Close()
		defer px.Close()
		go func() { io.Copy(px, cl) }()
		go func() { io.Copy(cl, px) }()
	}
}

// Shutdown will shutdown the server gracefully.
func (s *Server) Shutdown() error {
	return s.ln.Close()
}
