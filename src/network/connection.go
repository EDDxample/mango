package network

import (
	"mango/src/logger"
	"mango/src/utils"
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
	incomingPackets chan *utils.Buffer
	outgoingPackets chan *utils.Buffer
}

func NewConnection(connection net.Conn) IConnection {
	instance := &Connection{
		connection:      connection,
		running:         true,
		incomingPackets: make(chan *utils.Buffer, 10),
		outgoingPackets: make(chan *utils.Buffer, 10),
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
				HandlePacket(c, packet.GetUnsignedArray())
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

		buffer := utils.NewBufferWith(data)

		packetLength := buffer.ReadVarInt()

		i := buffer.GetInputIndex()
		packetBuffer := utils.NewBufferWith(buffer.GetUnsignedArray()[i-1 : i+packetLength])

		// packetBuffer := utils.NewBufferWith(data)

		logger.Debug("[C > S] %d", packetLength)
		c.incomingPackets <- packetBuffer
	}
	c.running = false
}

// Consumes the `outgoingPackets` channel and sends the packets to the client
func (c *Connection) handleOutgoingPackets() {
	for c.running {
		select {
		case packet := <-c.outgoingPackets:
			arr := packet.GetUnsignedArray()
			logger.Debug("[S > C] %+v", arr)
			c.connection.Write(arr)
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
