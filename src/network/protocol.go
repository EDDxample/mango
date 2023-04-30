package network

import (
	"bytes"
	"io"
	"mango/src/config"
	"mango/src/network/packet"
	"mango/src/network/packet/c2s"
	"mango/src/network/packet/s2c"
)

func HandleHandshakePacket(conn *Connection, data *[]byte) {
	reader := bytes.NewReader(*data)

	var handshake c2s.Handshake
	handshake.ReadPacket(reader)
	conn.state = ConnectionState(handshake.NextState)
}

func HandleStatusPacket(conn *Connection, data *[]byte) {
	reader := bytes.NewReader(*data)

	var header packet.PacketHeader
	header.ReadHeader(reader)

	reader.Seek(0, io.SeekStart)

	switch header.PacketID {
	case 0x00: // status packet
		var statusRequest c2s.StatusRequest
		statusRequest.ReadPacket(reader)

		var statusResponse s2c.Status
		statusResponse.Header.PacketID = 0
		statusResponse.StatusData.Protocol = uint16(762) // conn.Protocol

		packetBytes := statusResponse.Bytes()
		conn.outgoingPackets <- &packetBytes

	case 0x01: // ping packet
		var ping c2s.PingRequest
		ping.ReadPacket(reader)

		var pong s2c.PingResponse
		pong.Header.PacketID = 1
		pong.Timestamp = ping.Timestamp

		packetBytes := pong.Bytes()
		conn.outgoingPackets <- &packetBytes
	}
}

func HandleLoginPacket(conn *Connection, data *[]byte) {
	reader := bytes.NewReader(*data)

	var header packet.PacketHeader
	header.ReadHeader(reader)

	reader.Seek(0, io.SeekStart)

	switch header.PacketID {
	case 0x00: // Login Start
		var loginStart c2s.LoginStart
		loginStart.ReadPacket(reader)

		if config.IsOnline() {
			// TODO: implement cypher and return EncryptionRequest

		} else { // Offline mode, return LoginSuccess
			var logingSuccess s2c.LoginSuccess
			logingSuccess.Header.PacketID = 2
			logingSuccess.Username = loginStart.Name
			if loginStart.HasUUID {
				logingSuccess.UUID = loginStart.UUID
			}

			packetBytes := logingSuccess.Bytes()
			conn.outgoingPackets <- &packetBytes
			conn.state = PLAY
		}
	}
}

func HandlePlayPacket(conn *Connection, data *[]byte) {
	reader := bytes.NewReader(*data)

	var header packet.PacketHeader
	header.ReadHeader(reader)

	reader.Seek(0, io.SeekStart)

	switch header.PacketID {
	case 0x28: // Login (Play)
	case 0x50: // SetDefaultSpawnPosition
	}
}
