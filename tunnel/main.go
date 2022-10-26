package main

import (
	"flag"
	"tunnel/client"
	"tunnel/server"
)

var (
	port     int
	addr     string
	cli, srv bool
)

func init() {
	flag.IntVar(&port, "port", 8080, "listen port")
	flag.StringVar(&addr, "addr", "", "proxy to address")
	flag.BoolVar(&cli, "client", false, "client")
	flag.BoolVar(&srv, "server", false, "server")
	flag.Parse()
}

func main() {
	if cli {
		client.Client(port, addr)
	}
	if srv {
		server.Server(port, addr)
	}
}
