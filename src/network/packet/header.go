package packet

import (
	"io"
	dt "mango/src/network/datatypes"
)

type PacketHeader struct {
	Length   dt.VarInt
	PacketID dt.VarInt
}

func (pb *PacketHeader) ReadHeader(reader io.Reader) {
	pb.Length.ReadFrom(reader)
	pb.PacketID.ReadFrom(reader)
}

func (pb *PacketHeader) WriteHeader(buffer *[]byte) {
	*buffer = append(pb.PacketID.Bytes(), *buffer...)
	pb.Length = dt.VarInt(len(*buffer))
	*buffer = append(pb.Length.Bytes(), *buffer...)
}
