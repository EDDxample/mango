package network

import (
	"mango/src/config"
	"mango/src/logger"
	"mango/src/network/packet/c2s"
	"mango/src/network/packet/s2c"
	"net"
)

const (
	PROTOCOL_STATUS = 1
	PROTOCOL_LOGIN  = 2
)

func Handshake(conn net.Conn) {
	var handshake c2s.Handshake
	handshake.ReadPacket(conn)

	switch handshake.NextState {
	case PROTOCOL_STATUS:
		logger.Info("Phase: STATUS")

		// status
		var request c2s.Request
		request.ReadPacket(conn)

		var status s2c.Status
		status.Header.PacketID = 0
		status.StatusData.Protocol = uint16(handshake.Protocol)
		status.WritePacket(conn)

		// ping
		var ping c2s.Ping
		ping.ReadPacket(conn)

		var pong s2c.Pong
		pong.Header.PacketID = 1
		pong.Timestamp = ping.Timestamp
		pong.WritePacket(conn)

	case PROTOCOL_LOGIN:
		logger.Info("Phase: LOGIN")

		var request c2s.LoginStart
		request.ReadPacket(conn)

		if config.GConfig().IsOnline() {
			// TODO implement cypher
		}

		var response s2c.LoginSuccess
		response.Username = request.Name
		response.WritePacket(conn)

		logger.Info("Phase: PLAY")
	}
}
