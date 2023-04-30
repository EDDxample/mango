package network

import (
	"bytes"
	"io"
	"mango/src/logger"
	dt "mango/src/network/datatypes"
	"net"
	"time"
)

type IConnection interface {
	Tick()
	Close()
}

type ConnectionState int

const (
	SHAKE ConnectionState = iota
	STATUS
	LOGIN
	PLAY
)

type Connection struct {
	connection      net.Conn
	running         bool
	state           ConnectionState
	incomingPackets chan *[]byte
	outgoingPackets chan *[]byte
}

func NewConnection(connection net.Conn) IConnection {
	instance := &Connection{
		connection:      connection,
		state:           SHAKE,
		running:         true,
		incomingPackets: make(chan *[]byte, 10),
		outgoingPackets: make(chan *[]byte, 10),
	}
	go instance.handleIncomingPackets()
	go instance.handleOutgoingPackets()
	return instance
}

func (c *Connection) Tick() {
	if c.running {
		for {
			select {
			case packet := <-c.incomingPackets:
				switch c.state {
				case SHAKE:
					HandleHandshakePacket(c, packet)
				case STATUS:
					HandleStatusPacket(c, packet)
				case LOGIN:
					HandleLoginPacket(c, packet)
				case PLAY:
					HandlePlayPacket(c, packet)
				}
			default:
				return
			}
		}
	}
}

// Listens for client packets and puts them in the `incomingPackets` channel
func (c *Connection) handleIncomingPackets() {
	defer c.connection.Close()
	var data []byte

	for c.running {

		data = make([]byte, 1024*4)
		c.setTimeout()
		size, err := c.connection.Read(data)

		// handle initial read errors
		if err != nil || size == 0 {
			logger.Info("Client disconnected: %s (Reason: %s)", c.connection.RemoteAddr(), err)
			break
		}

		// split packets and push them into incomingPackets
		reader := bytes.NewReader(data)
		for start := 0; start < size; {
			reader.Seek(int64(start), io.SeekStart)

			var packetLength dt.VarInt
			n, err := packetLength.ReadFrom(reader)
			if err != nil || packetLength == 0 {
				if err != nil {
					logger.Info("Client disconnected: %s (Reason: %s)", c.connection.RemoteAddr(), err)
					c.running = false
				}
				break
			}

			end := start + int(n) + int(packetLength)

			packetBytes := data[start:end]
			start = end

			logger.Debug("[S < %s] %d, %v", c.connection.RemoteAddr(), packetLength, packetBytes)
			c.incomingPackets <- &packetBytes
		}
	}
	c.running = false
}

// Consumes the `outgoingPackets` channel and sends the packets to the client
func (c *Connection) handleOutgoingPackets() {
	for c.running {
		select {
		case packet := <-c.outgoingPackets:
			logger.Debug("[S > %s] %+v", c.connection.RemoteAddr().String(), len(*packet))
			c.connection.Write(*packet)
		default:
			continue
		}
	}
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
	c.running = false
}
