package c2s

import (
	"io"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type Ping struct {
	Header     packet.PacketHeader
	Timestamp  dt.Long
}

func (pk *Ping) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
	pk.Timestamp.ReadFrom(reader)
}
