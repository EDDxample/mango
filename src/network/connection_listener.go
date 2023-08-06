package network

import (
	"fmt"
	"net"
	"sync"
)

type _IConnectionListener interface {
	Start(host string, port int) error
	Tick()
}

type ConnectionListener struct {
	connections []*Connection
	lock        sync.RWMutex
}

func NewConnectionListener() *ConnectionListener {
	return &ConnectionListener{}
}

func (listener *ConnectionListener) Start(host string, port int) error {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return fmt.Errorf("address resolution failed [%v]", err)
	}

	socket, err := net.ListenTCP("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to bind [%v]", err)
	}

	// start TCP listener thread
	go func() {
		for {
			connection, err := socket.AcceptTCP()
			if err != nil {
				break
			}
			listener.addConnection(connection)
		}
	}()

	return nil
}

func (listener *ConnectionListener) Tick() {
	listener.lock.RLock()
	defer listener.lock.RUnlock()

	i := 0
	for _, connection := range listener.connections {
		if connection.alive {
			connection.Tick()

			// move alive connections to the front
			listener.connections[i] = connection
			i++
		}
	}

	// clear and remove closing connections from the back
	// https://utcc.utoronto.ca/~cks/space/blog/programming/GoSlicesMemoryLeak
	for j := i; j < len(listener.connections); j++ {
		listener.connections[j] = nil
	}
	listener.connections = listener.connections[:i]
}

func (listener *ConnectionListener) addConnection(connection *net.TCPConn) {
	listener.lock.Lock()
	defer listener.lock.Unlock()

	connection.SetNoDelay(true)
	connection.SetKeepAlive(true)
	listener.connections = append(listener.connections, NewConnection(connection))
}
