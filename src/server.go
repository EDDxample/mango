package src

import (
	"fmt"
	"mango/src/config"
	"mango/src/helper"
	"mango/src/logger"
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
		logger.Fatal(err)
	}
	defer listener.Close()

	logger.Info("Listening on %s:%d...", address, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Error accepting the connection")
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Error while handling connection (origin: %s): %s\n", helper.GetPanicReportData(), err)
		}
	}()

	logger.Info("New connection from %s\n", conn.RemoteAddr())
	network.Handshake(conn)
}
