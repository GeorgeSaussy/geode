package test_server

import "net/http"

// HelloServer is just a type alias for a ServerWrapper
type HelloServer = ServerWrapper

// MakeHelloServer creates a
func MakeHelloServer(port int) *HelloServer {
	s := &http.Server{}
	s.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	return &ServerWrapper{
		Server: s,
		Port:   port,
	}
}
