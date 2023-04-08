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
	var dataBuf []byte

	if pk.UUID != nil && len(pk.UUID) != 0 {
		dataBuf = append(dataBuf, pk.UUID...)
	} else {
		uuid1 := dt.Long(0xEDD)
		uuid2 := dt.Long(0x1337)
		dataBuf = append(dataBuf, uuid1.Bytes()...)
		dataBuf = append(dataBuf, uuid2.Bytes()...)
	}

	dataBuf = append(dataBuf, pk.Username.Bytes()...)
	dataBuf = append(dataBuf, 0x00)
	pk.Header.WriteHeader(&dataBuf)

	return dataBuf
}
