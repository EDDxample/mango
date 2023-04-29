package utils

import (
	"encoding/binary"
	"io"
	"math"
)

type Buffer struct {
	inputIndex  int32
	outputIndex int32

	array []byte
}

var _ io.Reader = Buffer{}

func NewBuffer() *Buffer {
	return NewBufferWith(make([]byte, 0))
}

func NewBufferWith(data []byte) *Buffer {
	return &Buffer{
		inputIndex:  0,
		outputIndex: 0,
		array:       data,
	}
}

// internal ==============================

func (b *Buffer) readNext() byte {
	if b.inputIndex >= b.Len() {
		return 0
	}

	next := b.array[b.inputIndex]
	b.inputIndex++

	if b.outputIndex > 0 {
		b.outputIndex--
	}
	return next
}

func (b Buffer) Read(p []byte) (n int, err error) {
	readed := b.readMany(len(p))
	n = copy(p, readed)

	return n, nil
}

func (b *Buffer) readMany(count int) []byte {
	bytes := make([]byte, count)
	for i := 0; i < count; i++ {
		bytes[i] = b.readNext()
	}
	return bytes
}

func (b *Buffer) readVariable(maxBytes int) int64 {
	var i int
	var out int64

	for {
		currentByte := int64(b.readNext())
		out |= (currentByte & 0x7F) << (i * 7)

		if i++; i > maxBytes {
			panic("omg holy shit man wtf btw")
		}

		if currentByte&0x80 == 0 {
			break
		}
	}
	return out
}

// metadata ==============================

func (b *Buffer) Len() int32 {
	return int32(len(b.array))
}

func (b *Buffer) GetSignedArray() []int8 {
	return asSignedArray(b.array)
}

func (b *Buffer) GetUnsignedArray() []byte {
	return b.array
}

func (b *Buffer) GetInputIndex() int32 {
	return b.inputIndex
}

func (b *Buffer) GetOutputIndex() int32 {
	return b.outputIndex
}

func (b *Buffer) Seek(count int32) {
	b.inputIndex += count
}

func (b *Buffer) SeekEnd() {
	b.Seek(b.Len() - 1)
}

// read ==================================

func (b *Buffer) ReadBoolean() bool {
	return b.readNext() != 0
}

func (b *Buffer) ReadByte() (byte, error) {
	return b.readNext(), nil
}

func (b *Buffer) ReadUShort() uint16 {
	return uint16(b.readNext())<<8 | uint16(b.readNext())
}

func (b *Buffer) ReadShort() int16 {
	return int16(b.ReadUShort())
}

func (b *Buffer) ReadUInt() uint32 {
	return binary.BigEndian.Uint32(b.readMany(4))
}

func (b *Buffer) ReadInt() int32 {
	return int32(b.ReadUInt())
}

func (b *Buffer) ReadULong() uint64 {
	return binary.BigEndian.Uint64(b.readMany(8))
}

func (b *Buffer) ReadLong() int64 {
	return int64(b.ReadULong())
}

func (b *Buffer) ReadFloat() float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(b.readMany(4)))
}

func (b *Buffer) ReadDouble() float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(b.readMany(8)))
}

func (b *Buffer) ReadVarInt() int32 {
	return int32(b.readVariable(5))
}

func (b *Buffer) ReadVarLong() int64 {
	return b.readVariable(10)
}

func (b *Buffer) ReadTxt() string {
	return string(b.ReadUnsignedArray())
}
func (b *Buffer) ReadUnsignedArray() []byte {
	size := b.ReadVarInt()
	array := b.array[b.inputIndex : b.inputIndex+size]
	b.inputIndex += size

	return array
}

func (b *Buffer) ReadSignedArray() []int8 {
	return asSignedArray(b.ReadUnsignedArray())
}

// push ==================================

// helpers ===============================

func asSignedArray(array []byte) []int8 {
	out := make([]int8, 0)
	for _, b := range array {
		out = append(out, int8(b))
	}
	return out
}

func asUnsignedArray(array []int8) []byte {
	out := make([]byte, 0)
	for _, b := range array {
		out = append(out, byte(b))
	}
	return out
}
