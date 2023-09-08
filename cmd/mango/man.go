package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"mango/src/config"
	"mango/src/logger"
	mango "mango/src/server"
)

var (
	configPath   = flag.String("c", "config.yml", "-c /path/to/config_file.yml")
	profilerPort = flag.Int("p", 0, "-p <profiler port>")
)

func main() {
	flag.Parse()
	config.Parse(*configPath)

	// profiler
	if *profilerPort != 0 {
		go runProfiler()
	}

	server := mango.NewServer()
	server.Start()
}

func runProfiler() {
	logger.Info("Profiling server runing on 127.0.0.1:%d...", *profilerPort)
	logger.Info(http.ListenAndServe(fmt.Sprintf(":%d", *profilerPort), nil))
}
