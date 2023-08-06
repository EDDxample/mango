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
			onSuccessfulLogin(conn)
		}
	}
}

func onSuccessfulLogin(conn *Connection) {
	var loginPlay s2c.LoginPlay
	packetBytes := loginPlay.Bytes()
	conn.outgoingPackets <- &packetBytes

	var spawnPos s2c.SetDefaultSpawnPosition
	packetBytes1 := spawnPos.Bytes()
	conn.outgoingPackets <- &packetBytes1

}
