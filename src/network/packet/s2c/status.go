package s2c

import (
	"io"
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

func (pk *Status) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
	pk.JsonPayload.ReadFrom(reader)
}

func (pk *Status) WritePacket(writer io.Writer) {
	pk.JsonPayload = dt.String(pk.getStatusPayload())

	var data []byte
	data = append(data, pk.JsonPayload.Bytes()...)

	pk.Header.WriteHeader(&data)

	writer.Write(data)
}
