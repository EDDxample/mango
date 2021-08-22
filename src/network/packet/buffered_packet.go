package packet

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
)

type BufferedPacket struct {
	Reader *bufio.Reader
	buffer []byte
}

// READ

func (bufferedPacket *BufferedPacket) ReadPacket(packet interface{}) interface{} {
	rType := reflect.TypeOf(packet)
	packetData := reflect.New(rType)

	for i := 0; i < rType.NumField(); i++ {
		field := packetData.Elem().Field(i)
		tags := rType.Field(i).Tag

		dataType := tags.Get("packet_dt")
		switch dataType {

		case "header":
			packetLength := bufferedPacket.ReadVarInt()
			packetID := bufferedPacket.ReadVarInt()
			field.Set(reflect.ValueOf(PacketHeader{packetLength, packetID}))

		case "varint":
			value := bufferedPacket.ReadVarInt()
			field.SetInt(int64(value))

		case "int":
			value := bufferedPacket.ReadInt()
			field.SetInt(int64(value))

		case "short":
			value := bufferedPacket.ReadUShort()
			field.SetInt(int64(value))

		case "long":
			value := bufferedPacket.ReadLong()
			field.SetInt(value)

		case "string":
			maxLength, _ := strconv.Atoi(tags.Get("max_length"))
			value := bufferedPacket.ReadString(int32(maxLength))
			field.SetString(value)

		}
	}
	reflect.ValueOf(&packet).Elem().Set(packetData.Elem())

	fmt.Printf("\nNew %s: %+v\n", rType.Name(), packet)

	return packet
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
func (bufferedPacket *BufferedPacket) ReadLong() int64 {
	var value int64
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

// WRITE

func (bufferedPacket *BufferedPacket) WriteBytes(data []byte) {
	bufferedPacket.buffer = append(bufferedPacket.buffer, data...)
}
func (bufferedPacket *BufferedPacket) WriteLong(value int64) {
	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, uint64(value))
	bufferedPacket.WriteBytes(buffer)
}
func (bufferedPacket *BufferedPacket) WriteVarInt(value int32) {
	bufferedPacket.WriteBytes(getVarInt(value))
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

// UTILS

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
