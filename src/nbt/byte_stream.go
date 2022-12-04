package nbt

import (
	"bytes"
	"encoding/binary"
	"math"
)

type NBTByteStream struct {
	stream *bytes.Reader
	output *bytes.Buffer
}

func NewNbtOutputStream() NBTByteStream {
	return NBTByteStream{
		output: new(bytes.Buffer),
	}
}

func (bs *NBTByteStream) ReadLongArray() ([]int64, error) {
	length, err := bs.ReadIntBE()
	if err != nil {
		return nil, err
	}

	arr := make([]int64, length)
	for i := int32(0); i < length; i++ {
		data, err := bs.ReadLongBE()
		if err != nil {
			return nil, err
		}

		arr[i] = data
	}

	return arr, nil
}

func (bs *NBTByteStream) ReadIntArray() ([]int32, error) {
	length, err := bs.ReadIntBE()
	if err != nil {
		return nil, err
	}

	arr := make([]int32, length)
	for i := int32(0); i < length; i++ {
		data, err := bs.ReadIntBE()
		if err != nil {
			return nil, err
		}

		arr[i] = data
	}

	return arr, nil
}

func (bs *NBTByteStream) ReadByteArray() ([]byte, error) {
	length, err := bs.ReadIntBE()
	if err != nil {
		return nil, err
	}

	arr := make([]byte, length)
	for i := int32(0); i < length; i++ {
		data, err := bs.ReadByte()
		if err != nil {
			return nil, err
		}

		arr[i] = data
	}

	return arr, nil
}

func (bs *NBTByteStream) ReadFloatBE() (float32, error) {
	var num uint32
	err := binary.Read(bs.stream, binary.BigEndian, &num)
	return math.Float32frombits(num), err
}

func (bs *NBTByteStream) ReadDoubleBE() (float64, error) {
	var num uint64
	err := binary.Read(bs.stream, binary.BigEndian, &num)

	return math.Float64frombits(num), err
}

func (bs *NBTByteStream) ReadString() (string, error) {
	strLength, err := bs.ReadShortBE()
	if err != nil {
		return "", err
	}

	name := make([]byte, strLength)
	_, err = bs.stream.Read(name)
	return string(name), err
}

func (bs *NBTByteStream) ReadByte() (byte, error) {
	return bs.stream.ReadByte()
}

func (bs *NBTByteStream) ReadShortBE() (int16, error) {
	var num int16
	err := binary.Read(bs.stream, binary.BigEndian, &num)
	return num, err
}

func (bs *NBTByteStream) ReadIntBE() (int32, error) {
	var num int32
	err := binary.Read(bs.stream, binary.BigEndian, &num)
	return num, err
}

func (bs *NBTByteStream) ReadLongBE() (int64, error) {
	var num int64
	err := binary.Read(bs.stream, binary.BigEndian, &num)
	return num, err
}

func (bs *NBTByteStream) WriteByte(num byte) {
	bs.output.WriteByte(num)
}

func (bs *NBTByteStream) WriteShort(num int16) {
	binary.Write(bs.output, binary.BigEndian, num)
}

func (bs *NBTByteStream) WriteInt(num int32) {
	binary.Write(bs.output, binary.BigEndian, &num)
}

func (bs *NBTByteStream) WriteLong(num int64) {
	binary.Write(bs.output, binary.BigEndian, &num)
}

func (bs *NBTByteStream) WriteFloat(num float32) {
	binary.Write(bs.output, binary.BigEndian, math.Float32bits(num))
}

func (bs *NBTByteStream) WriteDouble(num float64) {
	binary.Write(bs.output, binary.BigEndian, math.Float64bits(num))
}

func (bs *NBTByteStream) WriteString(str string) {
	bs.WriteShort(int16(len(str)))
	bs.output.WriteString(str)
}

func (bs *NBTByteStream) WriteByteArray(barr []byte) {
	bs.WriteInt(int32(len(barr)))
	bs.output.Write(barr)
}

func (bs *NBTByteStream) WriteIntArray(iarr []int32) {
	bs.WriteInt(int32(len(iarr)))

	for _, val := range iarr {
		bs.WriteInt(val)
	}
}

func (bs *NBTByteStream) WriteLongArray(larr []int64) {
	bs.WriteInt(int32(len(larr)))

	for _, val := range larr {
		bs.WriteLong(val)
	}
}
