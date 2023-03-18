package network

import (
	"bytes"
	"fmt"
	"mango/src/config"
	"mango/src/logger"
	"mango/src/network/packet/c2s"
	"mango/src/network/packet/s2c"
	"mango/src/utils"
)

const (
	PROTOCOL_PING               = 0x00
	PROTOCOL_STATUS             = 0x01
	PROTOCOL_LOGIN              = 0x02
	PROTOCOL_CLIENT_INFORMATION = 0x07
)

func protocolToString(protocol int32) string {
	switch protocol {
	case PROTOCOL_PING:
		return "PROTOCOL_PING"
	case PROTOCOL_STATUS:
		return "PROTOCOL_STATUS"
	case PROTOCOL_LOGIN:
		return "PROTOCOL_LOGIN"
	case PROTOCOL_CLIENT_INFORMATION:
		return "PROTOCOL_CLIENT_INFORMATION"
	default:
		return fmt.Sprintf("UNKWONW: %d", protocol)
	}
}

func HandlePacket(conn *Connection, data []byte) {
	logger.Debug("Ini: HandlePacket")

	reader := bytes.NewReader(data)

	var handshake c2s.Handshake
	handshake.ReadPacket(reader)
	logger.Debug("Handshake protocol from %s, next state: %s", conn.connection.RemoteAddr().String(), protocolToString(int32(handshake.NextState)))
	logger.Debug("Packet readed: %+v", handshake)

	switch handshake.NextState {
	case PROTOCOL_PING:
		logger.Debug("Phase PING")
		var ping c2s.Ping
		ping.ReadPacket(reader)

		var pong s2c.Pong
		pong.Header.PacketID = 1
		pong.Timestamp = ping.Timestamp
		conn.outgoingPackets <- utils.NewBufferWith(pong.Bytes())
		logger.Debug("PING Ended")
	case PROTOCOL_STATUS:
		logger.Debug("Phase STATUS")

		// Status
		var request c2s.Request
		request.ReadPacket(reader)

		var status s2c.Status
		status.Header.PacketID = 0
		status.StatusData.Protocol = uint16(handshake.Protocol)
		conn.outgoingPackets <- utils.NewBufferWith(status.Bytes())
		logger.Debug("STATUSU Ended")
	case PROTOCOL_LOGIN:
		logger.Debug("Phase: LOGIN")

		var request c2s.LoginStart
		request.ReadPacket(reader)

		if config.GConfig().IsOnline() {
			// TODO implement cypher
		}

		var response s2c.LoginSuccess
		response.Username = request.Name
		conn.outgoingPackets <- utils.NewBufferWith(response.Bytes())
		logger.Debug("LOGIN Ended")
	case PROTOCOL_CLIENT_INFORMATION:
		logger.Debug("Phase: CLIENT_INFORMATION")

		var request c2s.ClientInformation
		request.ReadPacket(reader)

		logger.Debug("ClientInformation: %+v", request)
	}
}
