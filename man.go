package main

import (
	"log"
	"mango/src/network"
)

const (
	HOST     = "localhost"
	PORT     = "25565"
	PROTOCOL = "tcp"
)

func main() {
	log.Println("Running...")
	network.Run(HOST, PORT, PROTOCOL)
}
