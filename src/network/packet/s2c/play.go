package s2c

import (
	dt "mango/src/network/datatypes"
)

type Play struct {
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

func Bytes() []byte {
	return nil
}
