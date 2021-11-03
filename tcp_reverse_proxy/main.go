package main

import (
	"flag"
	tpf "geode/tcp_reverse_proxy/tcp_reverse_proxy"
	"log"
)

func main() {
	lp := flag.Int("listen_port", 0, "the port on which the server will listen")
	pp := flag.Int("proxy_port", 0, "the port the forward will send proxied connections")
	flag.Parse()

	s, err := tpf.NewServer(*lp, *pp)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
	log.Println("server exited gracefully")
}
