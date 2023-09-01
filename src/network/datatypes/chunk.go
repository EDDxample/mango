package datatypes

type ChunkSection struct {
	NonAirBlocks Short
	BlockStates  PalettedContainer
	Biomes       PalettedContainer
}

func (sc *ChunkSection) Bytes() (buffer []byte) {
	buffer = append(buffer, sc.NonAirBlocks.Bytes()...)
	buffer = append(buffer, sc.BlockStates.Bytes()...)
	buffer = append(buffer, sc.Biomes.Bytes()...)
	return
}

// =====================================================
type PalettedContainer struct {
	BytesPerEntry UByte
	Palette       []VarInt
	Data          []Long
}

func (pc *PalettedContainer) Bytes() (buffer []byte) {
	buffer = append(buffer, pc.BytesPerEntry.Bytes()...)

	// single value palette
	if pc.BytesPerEntry == 0 {
		buffer = append(buffer, pc.Palette[0].Bytes()...)

	} else if pc.BytesPerEntry <= 8 { // indirect palette
		length := VarInt(len(pc.Palette))
		buffer = append(buffer, length.Bytes()...)

		for _, v := range pc.Palette {
			buffer = append(buffer, v.Bytes()...)
		}

	} // else { direct palette (no data) }

	// long array
	dataLength := VarInt(len(pc.Data))
	buffer = append(buffer, dataLength.Bytes()...)
	arr := make([]byte, 0, dataLength*8)
	for _, l := range pc.Data {
		arr = append(arr, l.Bytes()...)
	}
	buffer = append(buffer, arr...)

	return
}
