package network

import (
	"bufio"
	"encoding/binary"
)

func ReadVarInt(bufferReader *bufio.Reader) int32 {
	bitOffset := 0
	var value int32 = 0

	for {
		if bitOffset == 35 {
			panic("VarInt is too big")
		}

		currentByte, err := bufferReader.ReadByte()
		checkError(err)

		value |= int32(currentByte&0b01111111) << bitOffset

		bitOffset += 7

		if (currentByte & 0b10000000) == 0 {
			break
		}
	}

	return value
}

func ReadUShort(bufferReader *bufio.Reader) uint16 {
	var value uint16
	binary.Read(bufferReader, binary.BigEndian, &value)
	return value
}

func ReadInt(bufferReader *bufio.Reader) int32 {
	var value int32
	binary.Read(bufferReader, binary.BigEndian, &value)
	return value
}

func ReadString(bufferReader *bufio.Reader, maxLength int32) string {
	length := ReadVarInt(bufferReader)
	println("  - String Length:", length)
	if length > maxLength*4 {
		panic("String length cannot be larger than " + string(maxLength*4))
	}
	value := make([]byte, length)
	bufferReader.Read(value)
	return string(value)
}

func WriteVarInt(value int32) []byte {
	var buffer []byte

	for {
		currentByte := byte(value & 0b01111111)
		value >>= 7

		if value != 0 {
			currentByte |= 0b10000000
		}

		buffer = append(buffer, currentByte)

		if value == 0 {
			break
		}
	}
	return buffer
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
