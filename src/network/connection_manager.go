package network

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"mango/src/network/packet"
)

// listen for new connections
// handle handshake async
// if nextstate != status, add connection to array

func Run(address, port, protocol string) {
	socket, err := net.Listen(protocol, address+":"+port)
	if err != nil {
		log.Fatalf("Error while opening the connection socket, %s", err)
	}
	defer socket.Close()

	for {
		connection, err := socket.Accept()
		if err != nil {
			log.Printf("Couldn't accept an incoming socket, %s\n", err)
		} else {
			go HandleConnection(connection)
		}
	}
}

func HandleConnection(connection net.Conn) {
	var err error
	fmt.Printf("\nNew connection %s---------\n", connection.RemoteAddr())

	bufferedPacket := packet.BufferedPacket{Reader: bufio.NewReader(connection)}

	handshakePacket := bufferedPacket.ReadPacket(packet.C2SHandshake{}).(packet.C2SHandshake)

	switch handshakePacket.NextState {
	case STATUS:
		bufferedPacket.ReadPacket(packet.C2SRequest{}) // no need for casting if you're not using it
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
			uuid, err = getUUID(loginPacket.Username)
			if err != nil {
				log.Printf("Couldn't find UUID for username %s, %s\n", loginPacket.Username, err)
				connection.Close()
				return
			}
		}

		packet.WriteS2CLoginSuccess(connection, loginPacket.Username, uuid)
		handshakePacket.NextState = PLAY

	case PLAY:
	}
}

type UserUUID struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func getUUID(userName string) (string, error) {
	res, err := http.Get(fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", userName))
	if err != nil {
		return "", err
	}

	var userUuid UserUUID
	err = json.NewDecoder(res.Body).Decode(&userUuid)
	if err != nil {
		return "", err
	}
	log.Printf("Username %s corresponds to id %s and user %s", userName, userUuid.Id, userUuid.Name)

	return userUuid.Id, nil
}
