package client

import (
	"fmt"
	"net"
	"tunnel/lib"
)

func Client(port int, addr string) {
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
	defer conn.Close()
	server, err := net.Dial("tcp", addr)
	defer server.Close()
	if err != nil {
		fmt.Printf("server dial error: %v\n", err)
		return
	}
	go lib.Encode(conn, server)
	lib.Decode(server, conn)
}
