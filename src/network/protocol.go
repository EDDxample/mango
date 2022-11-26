package network

import (
	"log"
	"mango/src/network/packet/c2s"
	"mango/src/network/packet/s2c"
	"net"
)

const (
	PROTOCOL_STATUS = 1
)

func Handshake(conn net.Conn) {
	var handshake c2s.Handshake
	handshake.ReadPacket(conn)

	switch handshake.NextState {
	case PROTOCOL_STATUS:
		log.Println("[i] Type: STATUS")

		// status
		var request c2s.Request
		request.ReadPacket(conn)

		var status s2c.Status
		status.Header.PacketID = 0
		status.StatusData.Description = "Powered by man.go"
		status.StatusData.Protocol = uint16(handshake.Protocol)
		status.WritePacket(conn)

		// ping
		var ping c2s.Ping
		ping.ReadPacket(conn)

		var pong s2c.Pong
		pong.Header.PacketID = 1
		pong.Timestamp = ping.Timestamp
		pong.WritePacket(conn)
	}
}
