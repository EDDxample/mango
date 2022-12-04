package c2s

import (
	"io"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type Handshake struct {
	Header    packet.PacketHeader
	Protocol  dt.VarInt
	Address   dt.String
	Port      dt.UShort
	NextState dt.VarInt
}

func (pk *Handshake) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
	pk.Protocol.ReadFrom(reader)
	pk.Address.ReadFrom(reader)
	pk.Port.ReadFrom(reader)
	pk.NextState.ReadFrom(reader)
}
