package main

import (
	mango "mango/src"
	"mango/src/config"
)

func main() {
	config.Parse()
	mango.Start()
}
