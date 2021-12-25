package s2c

import (
	"io"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type Pong struct {
	Header    packet.PacketHeader
	Timestamp dt.Long
}

func (pk *Pong) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
	pk.Timestamp.ReadFrom(reader)
}

func (pk *Pong) WritePacket(writer io.Writer) {
	var data []byte
	data = append(data, pk.Timestamp.Bytes()...)
	pk.Header.WriteHeader(&data)

	writer.Write(data)
}
