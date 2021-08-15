package network

import (
	"bufio"
	"fmt"
	"net"

	"mango/src/network/packet"
)

// listen for new connections
// handle handshake async
// if nextstate != status, add connection to array

func Init(address string, port string) {
	socket, _ := net.Listen("tcp", address+":"+port)
	defer socket.Close()

	for {
		connection, _ := socket.Accept()
		go HandleConnection(connection)
	}
}

func HandleConnection(connection net.Conn) {
	fmt.Println("\nNew connection", connection.RemoteAddr().String(), "---------")

	bufferedPacket := packet.BufferedPacket{Reader: bufio.NewReader(connection)}

	nextState := packet.ReadC2SHandshake(&bufferedPacket)

	switch nextState {
	case STATUS:
		packet.ReadC2SRequest(&bufferedPacket)
		packet.WriteS2CStatus(connection)
		timestamp := packet.ReadC2SPing(&bufferedPacket)
		packet.WriteS2CPong(connection, timestamp)
		connection.Close()
	case LOGIN:

	case GAME:
	}
}
