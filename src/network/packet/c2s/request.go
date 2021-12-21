package c2s

import (
	"io"
	"mango/src/network/packet"
)

type Request struct {
	Header  packet.PacketHeader
}

func (pk *Request) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
}