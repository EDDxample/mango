package c2s

import (
	"io"
	"mango/src/network/packet"
)

type StatusRequest struct {
	Header packet.PacketHeader
}

func (pk *StatusRequest) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
}
