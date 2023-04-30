package c2s

import (
	"io"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type PingRequest struct {
	Header    packet.PacketHeader
	Timestamp dt.Long
}

func (pk *PingRequest) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
	pk.Timestamp.ReadFrom(reader)
}
