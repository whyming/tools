package server

import (
	"fmt"
	"net"
	"tunnel/lib"
)

func Server(port int, addr string) {
	client, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := client.Accept()
		if err != nil {
			panic(err)
		}
		go proxyCon(conn, addr)
	}
}

func proxyCon(conn net.Conn, addr string) {
	server, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("server dial error: %v\n", err)
		return
	}
	go lib.Decode(conn, server)
	lib.Encode(server, conn)
}
