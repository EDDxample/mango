package config

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

var gconfig GlobalConfig

type GlobalConfig struct {
	Server ServerConfig `yaml:"server"`
}

func (gc GlobalConfig) Motd() string {
	return gc.Server.Motd
}

func (gc GlobalConfig) Host() string {
	return gc.Server.Host
}

func (gc GlobalConfig) Port() int {
	return gc.Server.Port
}

func (gc GlobalConfig) IsOnline() bool {
	return gc.Server.Online
}

func (gc GlobalConfig) DebugMode() bool {
	return gc.Server.Debug
}

type ServerConfig struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Online bool   `yaml:"online"`
	Motd   string `yaml:"motd"`
	Debug  bool   `yaml:"debug"`
}

func Parse(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	yaml.Unmarshal([]byte(file), &gconfig)
}

func GConfig() GlobalConfig {
	return gconfig
}
