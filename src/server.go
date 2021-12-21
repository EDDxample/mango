package src

import (
	"fmt"
	"mango/src/network"
	"net"
)

const (
	address = "localhost"
	port = 25565
)

func Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil { panic(err) }
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil { panic(err) }
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("[+] New connection from %s\n", conn.RemoteAddr())
	network.Handshake(conn)
}