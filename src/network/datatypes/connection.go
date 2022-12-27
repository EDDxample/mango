package datatypes

import (
	"mango/src/logger"
	"net"
)

type RawPacket []byte

type Connection struct {
	RawConn          net.Conn
	InboundPackets   chan RawPacket
	OutboundPackets  chan RawPacket
	ShouldDisconnect bool
}

func NewConnection(conn net.Conn) *Connection {
	logger.Info("New connection from %s", conn.RemoteAddr().String())
	con := &Connection{
		RawConn:         conn,
		InboundPackets:  make(chan RawPacket),
		OutboundPackets: make(chan RawPacket, 10),
	}
	go con.ReadPackets()
	return con
}

// ReadPackets reads incomming packets and puts them into InboundPackets
// Called as a go routine
func (c *Connection) ReadPackets() {
	defer c.RawConn.Close()

	for !c.ShouldDisconnect {
		buffer := make([]byte, 1024*4)
		bytesRead, err := c.RawConn.Read(buffer)
		if err != nil || bytesRead == 0 {
			c.ShouldDisconnect = true
			break
		}

		c.InboundPackets <- buffer[:bytesRead]
	}
}
