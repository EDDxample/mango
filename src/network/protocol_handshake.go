package network

import (
	"bytes"
	"mango/src/logger"
	"mango/src/network/packet/c2s"
)

func HandleHandshakePacket(conn *Connection, data *[]byte) {
	reader := bytes.NewReader(*data)

	var handshake c2s.Handshake
	handshake.ReadPacket(reader)
	conn.state = Protocol(handshake.NextState)
	logger.Info("Handshake, next state -> %s", conn.state.ToString())
}
