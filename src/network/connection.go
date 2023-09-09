package network

import (
	"bytes"
	"io"
	"mango/src/logger"
	dt "mango/src/network/datatypes"
	"mango/src/network/packet/s2c"
	"net"
	"time"
)

type _IConnection interface {
	Tick()
	IsAlive() bool
	Close()
}

type Connection struct {
	connection      net.Conn
	alive           bool
	state           Protocol
	incomingPackets chan *[]byte
	outgoingPackets chan *[]byte
}

func NewConnection(connection net.Conn) *Connection {
	instance := &Connection{
		connection:      connection,
		state:           SHAKE,
		alive:           true,
		incomingPackets: make(chan *[]byte, 10),
		outgoingPackets: make(chan *[]byte, 10),
	}
	go instance.handleIncomingPackets()
	go instance.handleOutgoingPackets()
	return instance
}

func (c *Connection) Tick() {
	for c.alive {
		select {
		case packet := <-c.incomingPackets:
			HandlePacket(c, packet)

		default:
			return
		}
	}
}

// Listens for client packets and puts them in the `incomingPackets` channel
func (c *Connection) handleIncomingPackets() {
	defer c.Close()
	var data []byte

	for c.alive {

		data = make([]byte, 1024*4)
		c.setTimeout()
		size, err := c.connection.Read(data)

		// handle initial read errors
		if err != nil || size == 0 {
			logger.Info("Client disconnected: %s (Reason: %s)", c.connection.RemoteAddr(), err)
			break
		}

		// split packets and push them into `incomingPackets`
		reader := bytes.NewReader(data)
		for start := 0; start < size; {
			reader.Seek(int64(start), io.SeekStart)

			var packetLength dt.VarInt
			n, err := packetLength.ReadFrom(reader)
			if err != nil || packetLength == 0 {
				if err != nil {
					logger.Info("Client disconnected: %s (Reason: %s)", c.connection.RemoteAddr(), err)
					return
				}
				break
			}

			end := start + int(n) + int(packetLength)

			packetBytes := data[start:end]
			start = end

			// logger.Debug("[S < %s] %d, %v", c.connection.RemoteAddr(), packetLength, packetBytes)
			c.incomingPackets <- &packetBytes
		}
	}
}

// Consumes the `outgoingPackets` channel and sends the packets to the client
func (c *Connection) handleOutgoingPackets() {
	keepAliveTicker := time.NewTicker(10 * time.Second)
	var keepAlivePacket s2c.KeepAlive

	for c.alive {
		select {
		case packet := <-c.outgoingPackets:
			c.connection.Write(*packet)

		case t := <-keepAliveTicker.C:
			if c.state == PLAY {
				keepAlivePacket.KeepAliveID = dt.Long(t.UTC().UnixNano())
				c.connection.Write(keepAlivePacket.Bytes())
			}
		}
	}

	keepAliveTicker.Stop()
}

// Sets a different timeout depending on its current `state`
func (c *Connection) setTimeout() {
	if c.state == PLAY {
		c.connection.SetReadDeadline(time.Now().Add(10 * time.Second))
	} else {
		c.connection.SetReadDeadline(time.Now().Add(10 * time.Second))
	}
}

func (c *Connection) Close() {
	c.alive = false
	c.connection.Close()
}
