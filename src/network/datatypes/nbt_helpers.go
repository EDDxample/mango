package datatypes

import "mango/src/nbt"

func NBTCompound(data map[string]nbt.NBTTag) nbt.NBTTag {
	return nbt.NBTTag{NBTType: nbt.TAG_COMPOUND, NBTPayload: nbt.NBTPayloadUnion{Compound: data}}
}

func NBTString(text string) nbt.NBTTag {
	return nbt.NBTTag{NBTType: nbt.TAG_STRING, NBTPayload: nbt.NBTPayloadUnion{Str: text}}
}

func NBTByte(n byte) nbt.NBTTag {
	return nbt.NBTTag{NBTType: nbt.TAG_BYTE, NBTPayload: nbt.NBTPayloadUnion{S8: n}}
}

func NBTInt(n int32) nbt.NBTTag {
	return nbt.NBTTag{NBTType: nbt.TAG_INT, NBTPayload: nbt.NBTPayloadUnion{S32: n}}
}

func NBTLongArray(arr []int64) nbt.NBTTag {
	return nbt.NBTTag{NBTType: nbt.TAG_LONG_ARRAY, NBTPayload: nbt.NBTPayloadUnion{LongArray: arr}}
}

func NBTFloat(n float32) nbt.NBTTag {
	return nbt.NBTTag{NBTType: nbt.TAG_FLOAT, NBTPayload: nbt.NBTPayloadUnion{F32: n}}
}

func NBTDouble(n float64) nbt.NBTTag {
	return nbt.NBTTag{NBTType: nbt.TAG_DOUBLE, NBTPayload: nbt.NBTPayloadUnion{F64: n}}
}

func NBTList(list nbt.NBTList) nbt.NBTTag {
	return nbt.NBTTag{NBTType: nbt.TAG_LIST, NBTPayload: nbt.NBTPayloadUnion{List: list}}
}
