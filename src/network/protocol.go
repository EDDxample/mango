package network

import (
	"bytes"
	"mango/src/config"
	"mango/src/logger"
	"mango/src/network/datatypes"
	"mango/src/network/packet/c2s"
	"mango/src/network/packet/s2c"
)

const (
	PROTOCOL_PING               = 0x00
	PROTOCOL_STATUS             = 0x01
	PROTOCOL_LOGIN              = 0x02
	PROTOCOL_CLIENT_INFORMATION = 0x07
)

func HandlePacket(conn *datatypes.Connection, packet datatypes.RawPacket) {
	logger.Debug("Ini: HandlePacket")
	packetReader := bytes.NewReader(packet)

	var handshake c2s.Handshake
	handshake.ReadPacket(packetReader)
	logger.Debug("Handshake protocol from %s: %d", conn.RawConn.RemoteAddr().String(), handshake.NextState)

	switch handshake.NextState {
	case PROTOCOL_PING:
		var ping c2s.Ping
		ping.ReadPacket(packetReader)

		var pong s2c.Pong
		pong.Header.PacketID = 1
		pong.Timestamp = ping.Timestamp
		conn.OutboundPackets <- pong.Bytes()
		logger.Debug("PING Ended")
	case PROTOCOL_STATUS:
		logger.Info("Phase STATUS")

		// Status
		var request c2s.Request
		request.ReadPacket(packetReader)

		var status s2c.Status
		status.Header.PacketID = 0
		status.StatusData.Protocol = uint16(handshake.Protocol)
		conn.OutboundPackets <- datatypes.RawPacket(status.Bytes())
	case PROTOCOL_LOGIN:
		logger.Info("Phase: LOGIN")

		var request c2s.LoginStart
		request.ReadPacket(packetReader)

		if config.GConfig().IsOnline() {
			// TODO implement cypher
		}

		var response s2c.LoginSuccess
		response.Username = request.Name
		conn.OutboundPackets <- response.Bytes()
		logger.Debug("LOGIN Ended")
	case PROTOCOL_CLIENT_INFORMATION:
		logger.Info("Phase: CLIENT_INFORMATION")

		var request c2s.ClientInformation
		request.ReadPacket(packetReader)

		logger.Debug("ClientInformation: %+v", request)
	}
}
