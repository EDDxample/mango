package nbt

import "fmt"

// FIXME compress the output

func Marshal(data NBTTag) []byte {
	stream := NewNbtOutputStream()
	if data.NBTType == TAG_COMPOUND {
		stream.WriteByte(TAG_COMPOUND)
		stream.WriteString(data.Name)
		for k, v := range data.NBTPayload.Compound {
			stream.WriteByte(v.NBTType)
			stream.WriteString(k)
			write(v, &stream)
		}
		stream.WriteByte(TAG_END)
	}
	return stream.output.Bytes()
}

func write(data NBTTag, stream *NBTByteStream) {
	tag := data.NBTType
	switch tag {
	case TAG_BYTE:
		stream.WriteByte(data.NBTPayload.S8)
	case TAG_SHORT:
		stream.WriteShort(data.NBTPayload.S16)
	case TAG_INT:
		stream.WriteInt(data.NBTPayload.S32)
	case TAG_LONG:
		stream.WriteLong(data.NBTPayload.S64)
	case TAG_FLOAT:
		stream.WriteFloat(data.NBTPayload.F32)
	case TAG_DOUBLE:
		stream.WriteDouble(data.NBTPayload.F64)
	case TAG_BYTE_ARRAY:
		stream.WriteByteArray(data.NBTPayload.ByteArray)
	case TAG_STRING:
		stream.WriteString(data.NBTPayload.Str)
	case TAG_LIST:
		writeList(data.NBTPayload.List, stream)
	case TAG_COMPOUND:
		for k, v := range data.NBTPayload.Compound {
			stream.WriteByte(v.NBTType)
			stream.WriteString(k)
			write(v, stream)
		}
		stream.WriteByte(TAG_END)
	case TAG_INT_ARRAY:
		stream.WriteIntArray(data.NBTPayload.IntArray)
	case TAG_LONG_ARRAY:
		stream.WriteLongArray(data.NBTPayload.LongArray)
	default:
		panic(fmt.Sprintf("UNKNOWN TAG! %d", tag))
	}
}

func writeList(list NBTList, stream *NBTByteStream) {
	tag := list.Type
	stream.WriteByte(tag)

	switch tag {
	case TAG_BYTE:
		stream.WriteByteArray(list.Data.S8)
	case TAG_SHORT:
		length := int32(len(list.Data.S16))
		stream.WriteInt(length)
		for i := int32(0); i < length; i++ {
			stream.WriteShort(list.Data.S16[i])
		}
	case TAG_INT:
		stream.WriteIntArray(list.Data.S32)
	case TAG_LONG:
		stream.WriteLongArray(list.Data.S64)
	case TAG_FLOAT:
		length := int32(len(list.Data.F32))
		stream.WriteInt(length)
		for i := int32(0); i < length; i++ {
			stream.WriteFloat(list.Data.F32[i])
		}
	case TAG_DOUBLE:
		length := int32(len(list.Data.F64))
		stream.WriteInt(length)
		for i := int32(0); i < length; i++ {
			stream.WriteDouble(list.Data.F64[i])
		}
	case TAG_BYTE_ARRAY:
		length := int32(len(list.Data.ByteArray))
		stream.WriteInt(length)
		for i := int32(0); i < length; i++ {
			stream.WriteByteArray(list.Data.ByteArray[i])
		}
	case TAG_STRING:
		length := int32(len(list.Data.Str))
		stream.WriteInt(length)
		for i := int32(0); i < length; i++ {
			stream.WriteString(list.Data.Str[i])
		}
	case TAG_LIST:
		length := int32(len(list.Data.List))
		stream.WriteInt(length)
		for i := int32(0); i < length; i++ {
			writeList(list.Data.List[i], stream)
		}
	case TAG_COMPOUND:
		length := int32(len(list.Data.Compound))
		stream.WriteInt(length)
		for i := int32(0); i < length; i++ {
			for k, v := range list.Data.Compound[i] {
				stream.WriteByte(v.NBTType)
				stream.WriteString(k)
				write(v, stream)
			}
			stream.WriteByte(TAG_END)
		}
	case TAG_INT_ARRAY:
		length := int32(len(list.Data.IntArray))
		stream.WriteInt(length)
		for i := int32(0); i < length; i++ {
			stream.WriteIntArray(list.Data.IntArray[i])
		}
	case TAG_LONG_ARRAY:
		length := int32(len(list.Data.LongArray))
		stream.WriteInt(length)
		for i := int32(0); i < length; i++ {
			stream.WriteLongArray(list.Data.LongArray[i])
		}
	}
}
