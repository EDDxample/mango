package s2c

import (
	"mango/src/logger"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
	"os"
)

type LoginPlay struct {
	Header packet.PacketHeader
	// EntityID dt.Int
	IsHardcore          dt.Boolean
	Gamemode            dt.UByte
	PreviousGamemode    dt.Byte
	DimensionCount      dt.VarInt
	DimensionNames      []dt.String // Identifier
	RegistryCodec       dt.NbtCompound
	DimensionType       dt.String // Identifier
	DimensionName       dt.String // Identifier
	HashedSeed          dt.Long
	MaxPlayers          dt.VarInt
	ViewDistance        dt.VarInt
	SimulationDistance  dt.VarInt
	ReducedDebugInfo    dt.Boolean
	EnableRespawnScreen dt.Boolean
	IsDebug             dt.Boolean
	IsFlat              dt.Boolean
	HasDeathLocation    dt.Boolean
	DeathDimensionName  dt.String // Identifier
	DeathLocation       dt.String // FIXME: Position
}

func (pk *LoginPlay) Bytes() []byte {
	f, _ := os.Open("login_packet1.bin")
	arr := make([]byte, 30000)
	n, _ := f.Read(arr)
	logger.Info("sending %d bytes", n)
	return arr[:n]
}
