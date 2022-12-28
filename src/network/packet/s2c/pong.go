package s2c

import (
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type Pong struct {
	Header    packet.PacketHeader
	Timestamp dt.Long
}

func (pk *Pong) Bytes() []byte {
	var data []byte
	data = append(data, pk.Timestamp.Bytes()...)
	pk.Header.WriteHeader(&data)

	return data
}
