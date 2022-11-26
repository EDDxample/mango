package network

import (
	"log"
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
		log.Println("[i] Phase: STATUS")

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
		log.Println("[i] Phase: LOGIN")

		var request c2s.LoginStart
		request.ReadPacket(conn)

		var response s2c.LoginSuccess
		response.Username = request.Name
		response.WritePacket(conn)

		log.Println("[i] Phase: PLAY")
	}
}
