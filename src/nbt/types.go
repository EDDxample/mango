package nbt

import (
	"fmt"
	"strings"
)

var deep = 0

const (
	TAG_END byte = iota
	TAG_BYTE
	TAG_SHORT
	TAG_INT
	TAG_LONG
	TAG_FLOAT
	TAG_DOUBLE
	TAG_BYTE_ARRAY
	TAG_STRING
	TAG_LIST
	TAG_COMPOUND
	TAG_INT_ARRAY
	TAG_LONG_ARRAY
)

func TagTypeToString(tagType byte) string {
	switch tagType {
	case TAG_END:
		return "TAG_END"
	case TAG_BYTE:
		return "TAG_BYTE"
	case TAG_SHORT:
		return "TAG_SHORT"
	case TAG_INT:
		return "TAG_INT"
	case TAG_LONG:
		return "TAG_LONG"
	case TAG_FLOAT:
		return "TAG_FLOAT"
	case TAG_DOUBLE:
		return "TAG_DOUBLE"
	case TAG_BYTE_ARRAY:
		return "TAG_BYTE_ARRAY"
	case TAG_STRING:
		return "TAG_STRING"
	case TAG_LIST:
		return "TAG_LIST"
	case TAG_COMPOUND:
		return "TAG_COMPOUND"
	case TAG_INT_ARRAY:
		return "TAG_INT_ARRAY"
	case TAG_LONG_ARRAY:
		return "TAG_LONG_ARRAY"
	default:
		return "UNRECOGNIZED_TYPE"
	}
}

type NBTPayloadUnion struct {
	S8        byte
	S16       int16
	S32       int32
	S64       int64
	F32       float32
	F64       float64
	ByteArray []byte
	Str       string
	List      NBTList
	Compound  map[string]NBTTag
	IntArray  []int32
	LongArray []int64
}

type ListUnion struct {
	S8        []byte
	S16       []int16
	S32       []int32
	S64       []int64
	F32       []float32
	F64       []float64
	ByteArray [][]byte
	Str       []string
	List      []NBTList
	Compound  []map[string]NBTTag
	IntArray  [][]int32
	LongArray [][]int64
}

type NBTList struct {
	Type byte
	Data ListUnion
}

func (list NBTList) String() string {
	switch list.Type {
	case TAG_BYTE:
		return fmt.Sprintf("%v", list.Data.S8)
	case TAG_SHORT:
		return fmt.Sprintf("%v", list.Data.S16)
	case TAG_INT:
		return fmt.Sprintf("%v", list.Data.S32)
	case TAG_LONG:
		return fmt.Sprintf("%v", list.Data.S64)
	case TAG_FLOAT:
		return fmt.Sprintf("%v", list.Data.F32)
	case TAG_DOUBLE:
		return fmt.Sprintf("%v", list.Data.F32)
	case TAG_BYTE_ARRAY:
		return fmt.Sprintf("%v", list.Data.ByteArray)
	case TAG_STRING:
		return fmt.Sprintf("%v", list.Data.Str)
	case TAG_LIST:
		return list.String()
	case TAG_COMPOUND:
		return fmt.Sprintf("%s", list.Data.Compound)
	case TAG_INT_ARRAY:
		return fmt.Sprintf("%v", list.Data.IntArray)
	case TAG_LONG_ARRAY:
		return fmt.Sprintf("%v", list.Data.LongArray)
	default:
		return "UNKNOWN_ARRAY_TYPE"
	}
}

type NBTTag struct {
	NBTType    byte
	Name       string
	NBTPayload NBTPayloadUnion
}

func NewNamedCompound(name string) NBTTag {
	return NBTTag{
		NBTType: TAG_COMPOUND,
		Name:    name,
		NBTPayload: NBTPayloadUnion{
			Compound: make(map[string]NBTTag),
		},
	}
}

func NewCompound() NBTTag {
	return NBTTag{
		NBTType: TAG_COMPOUND,
		NBTPayload: NBTPayloadUnion{
			Compound: make(map[string]NBTTag),
		},
	}
}

func (tag NBTTag) GetByte(name string) byte {
	return tag.NBTPayload.Compound[name].NBTPayload.S8
}

func (tag NBTTag) SetByte(name string, data byte) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_BYTE,
		NBTPayload: NBTPayloadUnion{
			S8: data,
		},
	}
}

func (tag NBTTag) GetShort(name string) int16 {
	return tag.NBTPayload.Compound[name].NBTPayload.S16
}

func (tag NBTTag) SetShort(name string, data int16) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_SHORT,
		NBTPayload: NBTPayloadUnion{
			S16: data,
		},
	}
}

func (tag NBTTag) GetInt(name string) int32 {
	return tag.NBTPayload.Compound[name].NBTPayload.S32
}

func (tag NBTTag) SetInt(name string, data int32) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_INT,
		NBTPayload: NBTPayloadUnion{
			S32: data,
		},
	}
}

func (tag NBTTag) GetLong(name string) int64 {
	return tag.NBTPayload.Compound[name].NBTPayload.S64
}

func (tag NBTTag) SetLong(name string, data int64) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_LONG,
		NBTPayload: NBTPayloadUnion{
			S64: data,
		},
	}
}

func (tag NBTTag) GetFloat(name string) float32 {
	return tag.NBTPayload.Compound[name].NBTPayload.F32
}

func (tag NBTTag) SetFloat(name string, data float32) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_FLOAT,
		NBTPayload: NBTPayloadUnion{
			F32: data,
		},
	}
}

func (tag NBTTag) GetDouble(name string) float64 {
	return tag.NBTPayload.Compound[name].NBTPayload.F64
}

func (tag NBTTag) SetDouble(name string, data float64) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_DOUBLE,
		NBTPayload: NBTPayloadUnion{
			F64: data,
		},
	}
}

func (tag NBTTag) GetByteArray(name string) []byte {
	return tag.NBTPayload.Compound[name].NBTPayload.ByteArray
}

func (tag NBTTag) SetByteArray(name string, data []byte) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_BYTE_ARRAY,
		NBTPayload: NBTPayloadUnion{
			ByteArray: data,
		},
	}
}

func (tag NBTTag) GetString(name string) string {
	return tag.NBTPayload.Compound[name].NBTPayload.Str
}

func (tag NBTTag) SetString(name string, data string) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_STRING,
		NBTPayload: NBTPayloadUnion{
			Str: data,
		},
	}
}

func (tag NBTTag) GetList(name string) NBTList {
	return tag.NBTPayload.Compound[name].NBTPayload.List
}

func (tag NBTTag) SetList(name string, data NBTList) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_LIST,
		NBTPayload: NBTPayloadUnion{
			List: data,
		},
	}
}

func (tag NBTTag) GetCompound(name string) NBTTag {
	return tag.NBTPayload.Compound[name]
}

func (tag NBTTag) SetCompound(name string, data NBTTag) {
	tag.NBTPayload.Compound[name] = data
}

func (tag NBTTag) GetIntArray(name string) []int32 {
	return tag.NBTPayload.Compound[name].NBTPayload.IntArray
}

func (tag NBTTag) SetIntArray(name string, data []int32) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_INT_ARRAY,
		NBTPayload: NBTPayloadUnion{
			IntArray: data,
		},
	}
}

func (tag NBTTag) GetLongArray(name string) []int64 {
	return tag.NBTPayload.Compound[name].NBTPayload.LongArray
}

func (tag NBTTag) SetLongArray(name string, data []int64) {
	tag.NBTPayload.Compound[name] = NBTTag{
		NBTType: TAG_LONG_ARRAY,
		NBTPayload: NBTPayloadUnion{
			LongArray: data,
		},
	}
}

func (tag NBTTag) String() string {
	tabs := strings.Repeat("\t", deep)
	strFormat := tabs + "NBTTag(Type: %s, Name: %s, Payload: %s)\n"
	intFormat := tabs + "NBTTag(Type: %s, Name: %s, Payload: %d)\n"
	floatFormat := tabs + "NBTTag(Type: %s, Name: %s, Payload: %f)\n"
	listFormat := tabs + "NBTTag(Type: %s, Name: %s, Payload: %v)\n"
	compoundFormat := tabs + "NBTTag(Type: %s, Name: %s)\n%s"
	payload := tag.NBTPayload
	str := ""

	switch tag.NBTType {
	case TAG_BYTE:
		str += fmt.Sprintf(intFormat, TagTypeToString(TAG_BYTE), tag.Name, payload.S8)
	case TAG_SHORT:
		str += fmt.Sprintf(intFormat, TagTypeToString(TAG_SHORT), tag.Name, payload.S16)
	case TAG_INT:
		str += fmt.Sprintf(intFormat, TagTypeToString(TAG_INT), tag.Name, payload.S32)
	case TAG_LONG:
		str += fmt.Sprintf(intFormat, TagTypeToString(TAG_LONG), tag.Name, payload.S64)
	case TAG_FLOAT:
		str += fmt.Sprintf(floatFormat, TagTypeToString(TAG_FLOAT), tag.Name, payload.F32)
	case TAG_DOUBLE:
		str += fmt.Sprintf(floatFormat, TagTypeToString(TAG_DOUBLE), tag.Name, payload.F64)
	case TAG_BYTE_ARRAY:
		str += fmt.Sprintf(listFormat, TagTypeToString(TAG_BYTE_ARRAY), tag.Name, payload.ByteArray)
	case TAG_STRING:
		str += fmt.Sprintf(strFormat, TagTypeToString(TAG_STRING), tag.Name, payload.Str)
	case TAG_LIST:
		return fmt.Sprintf(listFormat, TagTypeToString(TAG_LIST), tag.Name, payload.List)
	case TAG_COMPOUND:
		result := ""
		deep += 1
		for k, v := range tag.NBTPayload.Compound {
			v.Name = k
			result += v.String()
		}

		deep -= 1
		str += fmt.Sprintf(compoundFormat, TagTypeToString(TAG_COMPOUND), tag.Name, result)
	case TAG_INT_ARRAY:
		str += fmt.Sprintf(listFormat, TagTypeToString(TAG_INT_ARRAY), tag.Name, payload.IntArray)
	case TAG_LONG_ARRAY:
		str += fmt.Sprintf(listFormat, TagTypeToString(TAG_LONG_ARRAY), tag.Name, payload.LongArray)
	default:
		panic("UNREACHEABLE")
	}

	return str
}
