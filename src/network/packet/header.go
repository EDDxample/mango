package packet

import (
	"io"
	dt "mango/src/network/datatypes"
)

type PacketHeader struct {
	length     dt.VarInt
	packetID   dt.VarInt
}

func (pb *PacketHeader) ReadHeader(reader io.Reader) {
	pb.length.ReadFrom(reader)
	pb.packetID.ReadFrom(reader)
}