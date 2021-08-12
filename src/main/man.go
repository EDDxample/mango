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
	"mango/src/log"
	"net"
	"os"
)

const (
	host     = "localhost"
	port     = "8080"
	protocol = "tcp"
)

func main() {
	if len(os.Args) != 2 {
		log.LOGGER.PrintError("Usage: go run . <client|server>")
	} else if os.Args[1] == "client" {
		clientConnect()
	} else if os.Args[1] == "server" {
		serverLoop()
	}
}

func clientConnect() {
	connection, err := net.Dial(protocol, host+":"+port)

	if err != nil {
		log.LOGGER.PrintError("Error connecting" + err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")

		input, _ := reader.ReadString('\n')

		connection.Write([]byte(input))

		message, _ := bufio.NewReader(connection).ReadString('\n')
		if message != "" {
			log.LOGGER.Print("Server relay: " + message)
		}

	}
}

func serverLoop() {
	log.LOGGER.PushContext("main")
	defer log.LOGGER.PopContext()

	log.LOGGER.Print("Starting " + protocol + " server on " + host + ":" + port + "...")

	socket, err := net.Listen(protocol, host+":"+port)
	defer socket.Close()

	if err != nil {
		log.LOGGER.PrintError("Error listening: " + err.Error())
	}

	for {
		connection, err := socket.Accept()

		if err != nil {
			log.LOGGER.PrintError("Error connecting: " + err.Error())
			return
		}

		log.LOGGER.Print("Client " + connection.RemoteAddr().String() + " connected.")

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	buffer, err := bufio.NewReader(connection).ReadBytes('\n')

	if err != nil {
		log.LOGGER.Print("Client " + connection.RemoteAddr().String() + " left.")
		connection.Close()
		return
	}

	log.LOGGER.Print("Client message: " + string(buffer[:len(buffer)-1]))
	connection.Write(buffer)
	handleConnection(connection)
}
