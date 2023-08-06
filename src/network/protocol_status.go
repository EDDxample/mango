package network

import (
	"bytes"
	"io"
	"mango/src/config"
	"mango/src/network/packet"
	"mango/src/network/packet/c2s"
	"mango/src/network/packet/s2c"
)

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
