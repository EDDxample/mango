package datatypes

import (
	"encoding/binary"
	"errors"
	"io"
)

func ReadByte(reader io.Reader) (value byte, err error) {
	err = binary.Read(reader, binary.BigEndian, &value)
	return
}

type Short int16 // ======================================================

func (s *Short) ReadFrom(reader io.Reader) (n int64, err error) {
	err = binary.Read(reader, binary.BigEndian, &s)
	return
}

func (s *Short) Bytes() (buffer []byte) {
	buffer = make([]byte, 8)
	binary.BigEndian.PutUint16(buffer, uint16(*s))
	return buffer
}

type UShort uint16 // ====================================================

func (s *UShort) ReadFrom(reader io.Reader) (n int64, err error) {
	err = binary.Read(reader, binary.BigEndian, s)
	return
}

func (s *UShort) Bytes() (buffer []byte) {
	buffer = make([]byte, 8)
	binary.BigEndian.PutUint16(buffer, uint16(*s))
	return buffer
}

type String string // ====================================================

func (s *String) ReadFrom(reader io.Reader) (n int64, err error) {
	var length VarInt
	nn, err := length.ReadFrom(reader)
	if err != nil {
		return nn, err
	}
	n += nn

	stringBytes := make([]byte, length)
	_, err = reader.Read(stringBytes)
	if err != nil {
		return n, err
	}
	n += int64(length)

	*s = String(stringBytes)
	return n, nil
}

func (s *String) Bytes() (buffer []byte) {

	strBytes := []byte(*s)
	length := VarInt(len(strBytes))

	buffer = append(buffer, length.Bytes()...)
	buffer = append(buffer, strBytes...)

	return buffer
}

type VarInt int32 // =====================================================

func (vi *VarInt) ReadFrom(reader io.Reader) (n int64, err error) {
	var value uint32

	for curr := byte(0x80); curr&0x80 != 0; n++ {
		if n > 5 {
			return n, errors.New("VarInt too big")
		}

		curr, err = ReadByte(reader)
		if err != nil {
			return n, err
		}

		value |= uint32(curr&0x7F) << (7 * n)
	}

	*vi = VarInt(value)
	return
}

func (vi *VarInt) Bytes() (buffer []byte) {
	value := *vi

	for i := 0; i < 5; i++ {
		var current byte = byte(value & 0x7F)
		value >>= 7

		if value > 0 {
			current |= 0x80
		}

		buffer = append(buffer, current)

		if value == 0 {
			return buffer
		}

	}
	return
}

type Long int64 // =====================================================

func (l *Long) ReadFrom(reader io.Reader) (n int64, err error) {
	err = binary.Read(reader, binary.BigEndian, l)
	return
}

func (l *Long) Bytes() (buffer []byte) {
	buffer = make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, uint64(*l))
	return buffer
}
