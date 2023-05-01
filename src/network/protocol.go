package network

import (
	"bytes"
	"io"
	"mango/src/config"
	"mango/src/logger"
	"mango/src/network/packet"
	"mango/src/network/packet/c2s"
	"mango/src/network/packet/s2c"
)

func HandleHandshakePacket(conn *Connection, data *[]byte) {
	reader := bytes.NewReader(*data)

	var handshake c2s.Handshake
	handshake.ReadPacket(reader)
	conn.state = ConnectionState(handshake.NextState)
	logger.Info("Handshake, next state -> %d", handshake.NextState)
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

		var statusResponse s2c.StatusResponse
		statusResponse.StatusData.Protocol = uint16(config.Protocol())

		packetBytes := statusResponse.Bytes()
		conn.outgoingPackets <- &packetBytes

	case 0x01: // ping packet
		var ping c2s.PingRequest
		ping.ReadPacket(reader)

		var pong s2c.PingResponse
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
			logingSuccess.Username = loginStart.Name
			if loginStart.HasUUID {
				logingSuccess.UUID = loginStart.UUID
			}

			packetBytes := logingSuccess.Bytes()
			conn.outgoingPackets <- &packetBytes
			logger.Debug("Login Success: %+v", logingSuccess)
			conn.state = PLAY

			// send init PLAY packets (Login (Play) + Set Default Spawn Position)

			var loginPlay s2c.LoginPlay
			packetBytes2 := loginPlay.Bytes()
			conn.outgoingPackets <- &packetBytes2

			var spawnPos s2c.SetDefaultSpawnPosition
			packetBytes3 := spawnPos.Bytes()
			conn.outgoingPackets <- &packetBytes3
		}
	}
}

func HandlePlayPacket(conn *Connection, data *[]byte) {
	reader := bytes.NewReader(*data)

	var header packet.PacketHeader
	header.ReadHeader(reader)

	logger.Info("PLAY packet ID: %d", header.PacketID)

	reader.Seek(0, io.SeekStart)

	switch header.PacketID {
	}
}
