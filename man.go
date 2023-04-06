package main

import (
	"flag"

	"mango/src/config"
	mango "mango/src/server"
)

var (
	configPath = flag.String("c", "config.yml", "-c /path/to/config_file.yml")
)

func main() {
	flag.Parse()
	config.Parse(*configPath)

	server := mango.NewServer()
	server.Start()
}
