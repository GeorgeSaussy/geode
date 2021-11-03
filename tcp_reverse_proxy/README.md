# README

The TCP Port Forwarder only exists to forward TCP traffic from a privileged port (e.g. port 80)
to a non-privileged port (e.g. port 5000). 

## Build

Build with `go build //sa/geode/tcp_reverse_proxy/main.go` or `make build` in this directory.

Test with `make test`. The tests will open TCP ports in the range 4000-4999.

## Run

Run TPF as `main --listen_port=XXX --proxy_port=YYY` where `XXX` is the port the server
will listen on and `YYY` is the port that traffic will be forwarded to.
Running `main --listen_port=80 --proxy_port=5000` will forward all traffic from port 80
to port 5000.