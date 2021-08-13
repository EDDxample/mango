package main

/*
-> = send packet
<- = wait for packet
*/
/*
Client:
	[CLOSED]
	-> SYN     (active open)
	[SYN-SENT]
	<- SYN+ACK
	-> ACK
	[ESTABLISHED] (program works)

*/
/*
Server:
	[CLOSED]
	[LISTEN]   (passive open)
	<- SYN
	-> SYN+ACK
	[SYN-RCVD]
	<- ACK
	[ESTABLISHED] (program works)
*/
/*
Either end can close the connection:
- If you close:
	[ESTABLISHED]
	-> FIN
	[FIN-WAIT-1]
	<- ACK
	[FIN-WAIT-2]
	<- FIN
	-> ACK
	[TIME-WAIT]
	[CLOSED] After 4 minutes

- If the other end closes:
	[ESTABLISHED]
	<- FIN
	-> ACK
	[CLOSE-WAIT]
	-> FIN
	[LAST-ACK]
	<- ACK
	[CLOSED]
*/
/*
CONCEPTOS:
	- AF_INET, AF_INET6: Address Family, IPv4/IPv6
	- SOCK_STREAM: Socket type, reliable byte stream (TCP)
	- SOCK_DGRAM:  Socket type, unreliable datagrams (UDP)
	- IPPROTO: Protocols, TCP/UDP
*/

import (
	"bufio"
	"fmt"
	"mango/src/network"
	"net"
)

const (
	host     = "localhost"
	port     = "25565"
	protocol = "tcp"
)

func main() {
	// serverLoop()
	fmt.Println("running...")
	fmt.Println()
	socket, _ := net.Listen("tcp", "localhost:25565")
	defer socket.Close()

	for {
		connection, _ := socket.Accept()
		fmt.Println("New connection", connection.RemoteAddr().String())
		fmt.Println()

		bufferReader := bufio.NewReader(connection)

		state := ProcessHandShakePacket(bufferReader)

		switch state {
		case 1:
			SendStatusPacket(connection, bufferReader)
			break
		case 2:
			Login()
			break

		}
	}
}

func ProcessHandShakePacket(bufferReader *bufio.Reader) int32 {
	println("HandShake Packet")
	// packet header
	packetLength := network.ReadVarInt(bufferReader)
	println("- Packet Length:", packetLength)
	packetID := network.ReadVarInt(bufferReader)
	println("- Packet ID:", packetID)
	// packet data
	protocolVersion := network.ReadVarInt(bufferReader)
	println("  - Protocol Version:", protocolVersion)
	serverAddress := network.ReadString(bufferReader, 255)
	println("  - Server Address:", serverAddress)
	serverPort := network.ReadUShort(bufferReader)
	println("  - Server Port:", serverPort)
	nextState := network.ReadVarInt(bufferReader)
	println("  - NextState:", nextState)
	println()
	return nextState
}

func SendStatusPacket(connection net.Conn, bufferReader *bufio.Reader) {
	println("Request Packet")
	// request packet
	packetLength := network.ReadVarInt(bufferReader)
	println("- Packet Length:", packetLength)
	packetID := network.ReadVarInt(bufferReader)
	println("- Packet ID:", packetID)
	println()

	statusPayload := []byte(`{
		"description": { "text": "Powered by man.go! \\o/" },
		"version": { "name": "1.16.5", "protocol": 754 },
		"players": { "max": 5, "online": -69 }
	}`)

	stringLength := len(statusPayload)
	bufferedStringLength := network.WriteVarInt(int32(stringLength))
	statusPacketLength := 1 + int32(len(bufferedStringLength)) + int32(stringLength)
	bufferedStatusPacketLength := network.WriteVarInt(statusPacketLength)
	statusPacketID := byte(0)

	var packet []byte
	packet = append(packet, bufferedStatusPacketLength...)
	packet = append(packet, statusPacketID)
	packet = append(packet, bufferedStringLength...)
	packet = append(packet, statusPayload...)

	connection.Write(packet)

}

func Login() {
}
