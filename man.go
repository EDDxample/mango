package main

import (
	"flag"
	mango "mango/src"
	"mango/src/config"
)

var (
	configPath = flag.String("c", "config.yml", "-c /path/to/config_file.yml")
)

func main() {
	flag.Parse()
	config.Parse(*configPath)
	//mango.Start()
	server := mango.NewServer()
	server.Start()
}
