package s2c

import (
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type Status struct {
	Header      packet.PacketHeader
	JsonPayload dt.String
	StatusData  StatusData
}

type StatusData struct {
	Protocol uint16
}

func (pk *Status) getStatusPayload() string {
	return dt.GetDemoServerStatus(int(pk.StatusData.Protocol))
}

func (pk *Status) Bytes() []byte {
	pk.JsonPayload = dt.String(pk.getStatusPayload())

	var data []byte
	data = append(data, pk.JsonPayload.Bytes()...)

	pk.Header.WriteHeader(&data)

	return data
}
