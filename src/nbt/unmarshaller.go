package nbt

import (
	"bytes"
	"errors"
	"log"
)

func ReadFile(filename string) NBTTag {
	file, err := DecompressFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(file)
	nbtByteStream := NBTByteStream{
		stream: r,
	}

	return Read(&nbtByteStream)
}

func ReadByteArray(data []byte) NBTTag {
	r := bytes.NewReader(data)
	nbtByteStream := NBTByteStream{
		stream: r,
	}

	return Read(&nbtByteStream)
}

func Read(stream *NBTByteStream) NBTTag {
	return readImplicitCompound(stream)
}

func readImplicitCompound(stream *NBTByteStream) NBTTag {
	tagType, err := stream.ReadByte()
	if err != nil {
		log.Fatalf("Couldn't read tagType, %s\n", err)
	}

	if tagType != TAG_COMPOUND {
		log.Fatalf("Expected TAG_COMPOUND but got %s\n", TagTypeToString(tagType))
	}

	str, err := stream.ReadString()
	if err != nil {
		log.Fatalf("Error reading string %s\n", err)
	}

	base := GetNbtCompound(stream)
	base.Name = str
	return base
}

func GetNbtCompound(stream *NBTByteStream) NBTTag {
	nbtTagCompound := NewCompound()
	nextTag, err := stream.ReadByte()
	if err != nil {
		log.Fatal("Error reading the name tag, ", err)
	}

	for nextTag != TAG_END {
		name, err := stream.ReadString()
		if err != nil {
			log.Fatalf("Couln't read string, %s\n", err)
		}

		err = readNbtTag(nextTag, stream, &nbtTagCompound, name)
		if err != nil {
			log.Fatalf("Couldn't read NBT %s\n", err)
		}

		nextTag, err = stream.ReadByte()
		if err != nil {
			log.Fatal("Couldn't read next tag, ", err)
		}
	}

	return nbtTagCompound
}

func readNbtTag(tag byte, stream *NBTByteStream, nbtTagCompound *NBTTag, name string) error {
	switch tag {
	case TAG_BYTE:
		data, err := stream.ReadByte()
		if err != nil {
			return err
		}

		nbtTagCompound.SetByte(name, data)
	case TAG_SHORT:
		data, err := stream.ReadShortBE()
		if err != nil {
			return err
		}

		nbtTagCompound.SetShort(name, data)
	case TAG_INT:
		data, err := stream.ReadIntBE()
		if err != nil {
			return err
		}

		nbtTagCompound.SetInt(name, data)
	case TAG_LONG:
		data, err := stream.ReadLongBE()
		if err != nil {
			return err
		}

		nbtTagCompound.SetLong(name, data)
	case TAG_FLOAT:
		data, err := stream.ReadFloatBE()
		if err != nil {
			return err
		}

		nbtTagCompound.SetFloat(name, data)
	case TAG_DOUBLE:
		data, err := stream.ReadDoubleBE()
		if err != nil {
			return err
		}

		nbtTagCompound.SetDouble(name, data)
	case TAG_BYTE_ARRAY:
		data, err := stream.ReadByteArray()
		if err != nil {
			return err
		}

		nbtTagCompound.SetByteArray(name, data)
	case TAG_STRING:
		data, err := stream.ReadString()
		if err != nil {
			return err
		}

		nbtTagCompound.SetString(name, data)
	case TAG_LIST:
		list, err := readList(stream)
		if err != nil {
			return err
		}

		nbtTagCompound.SetList(name, list)
	case TAG_COMPOUND:
		data := GetNbtCompound(stream)

		nbtTagCompound.SetCompound(name, data)
	case TAG_INT_ARRAY:
		data, err := stream.ReadIntArray()
		if err != nil {
			return err
		}

		nbtTagCompound.SetIntArray(name, data)
	case TAG_LONG_ARRAY:
		data, err := stream.ReadLongArray()
		if err != nil {
			return err
		}

		nbtTagCompound.SetLongArray(name, data)
	default:
		return errors.New("unknown tag type")
	}

	return nil
}

func readList(stream *NBTByteStream) (NBTList, error) {
	tag, err := stream.ReadByte()
	if err != nil {
		return NBTList{}, err
	}

	length, err := stream.ReadIntBE()
	if err != nil {
		return NBTList{}, err
	}

	return readListBody(tag, int(length), stream)
}

func readListBody(tag byte, length int, stream *NBTByteStream) (NBTList, error) {
	var list NBTList

	switch tag {
	case TAG_BYTE:
		payload := make([]byte, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadByte()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_BYTE
		list.Data.S8 = payload

	case TAG_SHORT:
		payload := make([]int16, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadShortBE()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_SHORT
		list.Data.S16 = payload

	case TAG_INT:
		payload := make([]int32, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadIntBE()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_INT
		list.Data.S32 = payload

	case TAG_LONG:
		payload := make([]int64, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadLongBE()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_LONG
		list.Data.S64 = payload

	case TAG_FLOAT:
		payload := make([]float32, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadFloatBE()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_FLOAT
		list.Data.F32 = payload

	case TAG_DOUBLE:
		payload := make([]float64, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadDoubleBE()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_DOUBLE
		list.Data.F64 = payload

	case TAG_BYTE_ARRAY:
		payload := make([][]byte, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadByteArray()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_BYTE_ARRAY
		list.Data.ByteArray = payload

	case TAG_STRING:
		payload := make([]string, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadString()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_STRING
		list.Data.Str = payload

	case TAG_LIST:
		payload := make([]NBTList, 0)
		for i := 0; i < length; i++ {
			data, err := readList(stream)
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_LIST
		list.Data.List = payload

	case TAG_COMPOUND:
		payload := make([]map[string]NBTTag, 0)
		for i := 0; i < length; i++ {
			data := GetNbtCompound(stream)
			payload = append(payload, data.NBTPayload.Compound)
		}

		list.Type = TAG_COMPOUND
		list.Data.Compound = payload

	case TAG_INT_ARRAY:
		payload := make([][]int32, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadIntArray()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_INT_ARRAY
		list.Data.IntArray = payload

	case TAG_LONG_ARRAY:
		payload := make([][]int64, 0)
		for i := 0; i < length; i++ {
			data, err := stream.ReadLongArray()
			if err != nil {
				return list, err
			}

			payload = append(payload, data)
		}

		list.Type = TAG_LONG_ARRAY
		list.Data.LongArray = payload

	}

	return list, nil
}
