package main

import (
	"fmt"
	"mango/src/network"
)

const (
	host     = "localhost"
	port     = "25565"
	protocol = "tcp"
)

func main() {
	fmt.Println("running...")
	network.Init("localhost", "25565")
}
