package s2c

import (
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type LoginSuccess struct {
	Header   packet.PacketHeader
	Username dt.String
	UUID     []byte
}

func (pk *LoginSuccess) Bytes() []byte {
	pk.Header.PacketID = 2
	var data []byte

	if pk.UUID != nil && len(pk.UUID) != 0 {
		data = append(data, pk.UUID...)
	} else {
		uuid1 := dt.Long(0xEDD)
		uuid2 := dt.Long(0x1337)
		data = append(data, uuid1.Bytes()...)
		data = append(data, uuid2.Bytes()...)
	}

	data = append(data, pk.Username.Bytes()...)
	data = append(data, 0x00)
	pk.Header.WriteHeader(&data)

	return data
}
