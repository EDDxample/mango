package packet

import (
	"bufio"
	"encoding/binary"
)

type BufferedPacket struct {
	Reader *bufio.Reader
	buffer []byte
}

func (bufferedPacket *BufferedPacket) ReadVarInt() int32 {
	bitOffset := 0
	var value int32 = 0

	for {
		if bitOffset == 35 {
			panic("VarInt is too big")
		}

		currentByte, err := bufferedPacket.Reader.ReadByte()
		checkError(err)

		value |= int32(currentByte&0b01111111) << bitOffset

		bitOffset += 7

		if (currentByte & 0b10000000) == 0 {
			break
		}
	}

	return value
}

func (bufferedPacket *BufferedPacket) ReadUShort() uint16 {
	var value uint16
	binary.Read(bufferedPacket.Reader, binary.BigEndian, &value)
	return value
}

func (bufferedPacket *BufferedPacket) ReadInt() int32 {
	var value int32
	binary.Read(bufferedPacket.Reader, binary.BigEndian, &value)
	return value
}

func (bufferedPacket *BufferedPacket) ReadString(maxLength int32) string {
	length := bufferedPacket.ReadVarInt()
	if length > maxLength*4 {
		panic("String length cannot be larger than " + string(maxLength*4))
	}
	value := make([]byte, length)
	bufferedPacket.Reader.Read(value)
	return string(value)
}

func (bufferedPacket *BufferedPacket) ReadPacketHeader() (int32, int32) {
	packetLength := bufferedPacket.ReadVarInt()
	packetID := bufferedPacket.ReadVarInt()
	return packetLength, packetID
}

func (bufferedPacket *BufferedPacket) WriteBytes(data []byte) {
	bufferedPacket.buffer = append(bufferedPacket.buffer, data...)
}

func (bufferedPacket *BufferedPacket) WriteVarInt(value int32) {
	bufferedPacket.buffer = append(bufferedPacket.buffer, getVarInt(value)...)
}

func (bufferedPacket *BufferedPacket) WriteString(text string) {
	payload := []byte(text)
	bufferedPacket.WriteVarInt(int32(len(payload)))
	bufferedPacket.WriteBytes(payload)
}

func (bufferedPacket *BufferedPacket) GetBufferWithHeader(packetID int32) []byte {
	packetIDbuffer := getVarInt(packetID)
	packetLength := int32(len(packetIDbuffer)) + int32(len(bufferedPacket.buffer))
	header := append(getVarInt(packetLength), packetIDbuffer...)
	return append(header, bufferedPacket.buffer...)
}

func getVarInt(value int32) []byte {
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
