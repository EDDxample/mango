package src

import (
	"fmt"
	"mango/src/config"
	"mango/src/logger"
	"mango/src/network"
	"mango/src/network/datatypes"
	"net"
	"sync"
)

type MinecraftServer struct {
	Connections []*datatypes.Connection
	Mutex       sync.Mutex
}

func NewServer() *MinecraftServer {
	return &MinecraftServer{
		Connections: make([]*datatypes.Connection, 0),
		Mutex:       sync.Mutex{},
	}
}

func (server *MinecraftServer) Start() {
	logger.Debug("Server started")
	go server.StartTCP()

	for {
		server.TickConnections()
	}
}

func (server *MinecraftServer) StartTCP() {
	address := config.GConfig().Host()
	port := config.GConfig().Port()

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		logger.Error("Error listening: %s", err)
		return
	}
	defer ln.Close()
	logger.Info("Listening on %s:%d...", address, port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error("Error accpeting connection: %s", err)
			continue
		}

		// add to thread-safe connection list
		server.Mutex.Lock()
		server.Connections = append(server.Connections, datatypes.NewConnection(conn))
		server.Mutex.Unlock()
	}
}

// TickConnections ticks every connection and handles inbound/outbound packets accordingly
func (server *MinecraftServer) TickConnections() {
	var closingConnections []int

	server.Mutex.Lock()
	for i, connection := range server.Connections {
		if connection.ShouldDisconnect {
			closingConnections = append(closingConnections, i)
			continue
		}

		// read packets until both channels are empty
		stop := false
		for !stop {
			select {
			// handle inbound packets
			case packet := <-connection.InboundPackets:
				logger.Info("[%s] Received packet", connection.RawConn.RemoteAddr().String())
				server.HandlePacket(connection, packet)

			// handle outbound packets
			case packet := <-connection.OutboundPackets:
				logger.Info("[server] Sending outgoing to %s", connection.RawConn.RemoteAddr().String())
				connection.RawConn.Write(packet)

			// no packets found
			default:
				stop = true
			}
		}

	}

	// close connections
	for i, j := range closingConnections {
		logger.Info("[%s] Disconnected", server.Connections[j-i].RawConn.RemoteAddr().String())
		server.Connections = append(server.Connections[:j-i], server.Connections[j-i+1:]...)
	}
	server.Mutex.Unlock()
}

func (server *MinecraftServer) HandlePacket(conn *datatypes.Connection, packet datatypes.RawPacket) {
	network.HandlePacket(conn, packet)
}
