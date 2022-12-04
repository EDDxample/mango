package src

import (
	"fmt"
	"log"
	"mango/src/config"
	"mango/src/helper"
	"mango/src/network"
	"net"
)

var (
	address = "0.0.0.0"
	port    = 25565
)

func Start() {
	address = config.GConfig().Host()
	port = config.GConfig().Port()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("[i] Listening on %s:%d...\n", address, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[!] Error accepting the connection\n")
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[!] Error while handling connection (origin: %s): %s\n", helper.GetPanicReportData(), err)
		}
	}()

	log.Printf("[+] New connection from %s\n", conn.RemoteAddr())
	network.Handshake(conn)
}
