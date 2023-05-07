package s2c

import (
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type KeepAlive struct {
	Header      packet.PacketHeader
	KeepAliveID dt.Long
}

func (pk *KeepAlive) Bytes() []byte {
	pk.Header.PacketID = 0x23
	var data []byte

	data = append(data, pk.KeepAliveID.Bytes()...)
	pk.Header.WriteHeader(&data)

	return data
}
