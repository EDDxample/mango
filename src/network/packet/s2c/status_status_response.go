package s2c

import (
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type StatusResponse struct {
	Header      packet.PacketHeader
	JsonPayload dt.String
	StatusData  StatusData
}

type StatusData struct {
	Protocol uint16
}

func (pk *StatusResponse) getStatusPayload() string {
	return dt.GetDemoServerStatus(int(pk.StatusData.Protocol))
}

func (pk *StatusResponse) Bytes() []byte {
	pk.Header.PacketID = 0x00
	pk.JsonPayload = dt.String(pk.getStatusPayload())

	var data []byte
	data = append(data, pk.JsonPayload.Bytes()...)

	pk.Header.WriteHeader(&data)

	return data
}
