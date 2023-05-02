package s2c

import (
	"io/ioutil"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
	"os"
)

type LoginPlay struct {
	Header              packet.PacketHeader
	EntityID            dt.Int
	IsHardcore          dt.Boolean
	Gamemode            dt.UByte
	PreviousGamemode    dt.Byte
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
	DeathLocation       dt.Position
}

func (pk *LoginPlay) Bytes() []byte {
	// return pk.getRegularBytes() // TODO: use this version at some point
	return pk.getStoredPacketBytes()
}

// loads the full packet bytes
func (pk LoginPlay) getStoredPacketBytes() []byte {
	f, err := os.Open("assets/fullLoginPacket.bin")
	arr, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	n := len(arr)
	length := dt.VarInt(n)
	return append(length.Bytes(), arr...)
}

// Loads only the registryCodec bytes
func (pk LoginPlay) getStoredRegistryBytes() []byte {
	f, err := os.Open("assets/registryCodec.bin")
	arr, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	n := len(arr)
	return arr[:n]
}

// TODO: use this version at some point
func (pk *LoginPlay) getRegularBytes() []byte {
	pk.populatePacket()
	pk.Header.PacketID = 0x28
	var data []byte
	data = append(data, pk.EntityID.Bytes()...)
	data = append(data, pk.IsHardcore.Bytes()...)
	data = append(data, pk.Gamemode.Bytes()...)
	data = append(data, pk.PreviousGamemode.Bytes()...)

	ndims := dt.VarInt(len(pk.DimensionNames))
	data = append(data, ndims.Bytes()...)
	for _, dim := range pk.DimensionNames {
		data = append(data, dim.Bytes()...)
	}

	// data = append(data, pk.RegistryCodec.Bytes()...)
	data = append(data, pk.getStoredRegistryBytes()...)

	data = append(data, pk.DimensionType.Bytes()...)
	data = append(data, pk.DimensionName.Bytes()...)

	data = append(data, pk.HashedSeed.Bytes()...)
	data = append(data, pk.MaxPlayers.Bytes()...)

	data = append(data, pk.ViewDistance.Bytes()...)
	data = append(data, pk.SimulationDistance.Bytes()...)

	data = append(data, pk.ReducedDebugInfo.Bytes()...)
	data = append(data, pk.EnableRespawnScreen.Bytes()...)
	data = append(data, pk.IsDebug.Bytes()...)
	data = append(data, pk.IsFlat.Bytes()...)
	data = append(data, pk.HasDeathLocation.Bytes()...)

	if pk.HasDeathLocation {
		data = append(data, pk.DeathDimensionName.Bytes()...)
		data = append(data, pk.DeathLocation.Bytes()...)
	}

	pk.Header.WriteHeader(&data)
	return data
}

func (pk *LoginPlay) populatePacket() {
	pk.EntityID = 1
	pk.IsHardcore = false
	pk.Gamemode = 1
	pk.PreviousGamemode = 0xFF // -1 aka undefined
	pk.DimensionNames = []dt.String{
		"minecraft:overworld",
		"minecraft:the_end",
		"minecraft:the_nether",
	}

	pk.RegistryCodec = dt.GetDemoRegistryCodec()

	pk.DimensionType = "minecraft:overworld"
	pk.DimensionName = "minecraft:overworld"

	pk.HashedSeed = 0xEDD1337DeadFace

	pk.MaxPlayers = 69
	pk.ViewDistance = 10
	pk.SimulationDistance = 10

	pk.ReducedDebugInfo = false
	pk.EnableRespawnScreen = true
	pk.IsDebug = false
	pk.IsFlat = true
	pk.HasDeathLocation = false
	pk.DeathDimensionName = ""
	pk.DeathLocation = dt.Position{X: 0, Y: 0, Z: 0}
}
