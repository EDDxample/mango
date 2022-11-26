package c2s

import (
	"io"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type LoginStart struct {
	Header packet.PacketHeader
	Name   dt.String
}

func (pk *LoginStart) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
	pk.Name.ReadFrom(reader)
}
