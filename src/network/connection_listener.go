package network

import (
	"fmt"
	"net"
	"sync"
)

type IConnectionListener interface {
	Start(host string, port int) error
	Tick()
}

type ConnectionListener struct {
	connections []IConnection
	lock        sync.RWMutex
}

func NewConnectionListener() IConnectionListener {
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

	for _, connection := range listener.connections {
		connection.Tick()
	}
}

func (listener *ConnectionListener) addConnection(connection *net.TCPConn) {
	listener.lock.Lock()
	defer listener.lock.Unlock()

	connection.SetNoDelay(true)
	connection.SetKeepAlive(true)
	listener.connections = append(listener.connections, NewConnection(connection))
}
