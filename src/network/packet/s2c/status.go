package s2c

import (
	"fmt"
	"io"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet"
)

type Status struct {
	Header      packet.PacketHeader
	JsonPayload dt.String
	status      StatusData
}

type StatusData struct {
	Description string
	Protocol    uint16
}

func (pk *Status) getStatusPayload() string {
	return fmt.Sprintf(`{
		"description": { "text": "%s" },
		"version": { "name": "mango", "protocol": %d },
		"players": { "max": 69, "online": -1 },
		"favicon": "data:image/png;base64,<data>"
	}`,
		pk.status.Description,
		pk.status.Protocol,
	)
}

func (pk *Status) ReadPacket(reader io.Reader) {
	pk.Header.ReadHeader(reader)
	pk.JsonPayload.ReadFrom(reader)
}

func (pk *Status) WritePacket(writer io.Writer) {
	pk.JsonPayload = dt.String(pk.getStatusPayload())
	// var buffer bytes.Buffer
	// pk.JsonPayload.WriteTo(writer)
	// pk.Header.WriteHeader(writer)
}
