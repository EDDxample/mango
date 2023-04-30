package c2s

import (
	"io"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type LoginStart struct {
	Header  packet.PacketHeader
	Name    dt.String
	HasUUID dt.Boolean
	UUID    []byte
}

func (pk *LoginStart) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
	pk.Name.ReadFrom(reader)

	pk.HasUUID.ReadFrom(reader)
	if pk.HasUUID {
		pk.UUID = make([]byte, 16)

		if _, err := reader.Read(pk.UUID); err != nil {
			panic(err)
		}
	}
}
