package s2c

import (
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type SetDefaultSpawnPosition struct {
	Header   packet.PacketHeader
	Location dt.Position
	Angle    dt.Float
}

func (pk *SetDefaultSpawnPosition) Bytes() []byte {
	pk.Header.PacketID = 0x51
	var data []byte

	data = append(data, pk.Location.Bytes()...)
	data = append(data, pk.Angle.Bytes()...)
	pk.Header.WriteHeader(&data)

	return data
}
