package network

import (
	"bytes"
	"io"
	"mango/src/network/packet"
)

func HandlePlayPacket(conn *Connection, data *[]byte) {
	reader := bytes.NewReader(*data)

	var header packet.PacketHeader
	header.ReadHeader(reader)

	// logger.Info("PLAY packet ID: %d", header.PacketID)

	reader.Seek(0, io.SeekStart)

	switch header.PacketID {
	}
}
