package nbt

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
)

func DecompressFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return DecompressByteArray(data)
}

const (
	GZIP = 0x1f
	ZLIB = 0x78
)

func DecompressByteArray(byteArray []byte) ([]byte, error) {
	compressType := byteArray[0]
	switch compressType {
	case GZIP:
		return decompressGzip(byteArray)
	case ZLIB:
		return decompressZlib(byteArray)
	default:
		return byteArray, nil
	}
}
func decompressZlib(file []byte) ([]byte, error) {
	panic("UNIMPLEMENTED")
}

func decompressGzip(file []byte) ([]byte, error) {
	gzipr, err := gzip.NewReader(bytes.NewBuffer(file))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(gzipr)
	if err != nil {
		return nil, err
	}
	gzipr.Close()

	return data, nil
}
