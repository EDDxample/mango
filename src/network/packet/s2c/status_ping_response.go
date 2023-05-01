package s2c

import (
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type PingResponse struct {
	Header    packet.PacketHeader
	Timestamp dt.Long
}

func (pk *PingResponse) Bytes() []byte {
	pk.Header.PacketID = 0x01
	var data []byte
	data = append(data, pk.Timestamp.Bytes()...)
	pk.Header.WriteHeader(&data)

	return data
}
