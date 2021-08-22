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

	handshakePacket := bufferedPacket.ReadPacket(packet.C2SHandshake{}).(packet.C2SHandshake)

	switch handshakePacket.NextState {
	case STATUS:
		_ = bufferedPacket.ReadPacket(packet.C2SRequest{}) // no need for casting if you're not using it
		packet.WriteS2CStatus(connection)
		pingPacket := bufferedPacket.ReadPacket(packet.C2SPing{}).(packet.C2SPing)
		packet.WriteS2CPong(connection, pingPacket.Timestamp)
		connection.Close()

	case LOGIN:
		premiumServer := false
		uuid := ""

		loginPacket := bufferedPacket.ReadPacket(packet.C2SLoginStart{}).(packet.C2SLoginStart)

		if premiumServer {
			// Client auth
			//   C→S: Encryption Response
			// Server auth, both enable encryption
			//   S→C: Set Compression (optional)
		} else {
			uuid = getUUID(loginPacket.Username)
		}

		packet.WriteS2CLoginSuccess(connection, loginPacket.Username, uuid)
		handshakePacket.NextState = PLAY

	case PLAY:
	}
}

func getUUID(userName string) string {
	// idk
	return "396367fa-b5d1-3a3f-b390-ea07a86c3112"
}
