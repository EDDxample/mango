package config

import (
	"log"
	"os"
	"strings"

	"github.com/go-yaml/yaml"
)

var gconfig GlobalConfig

type GlobalConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Logger   LoggerConfig   `yaml:"logger"`
	Profiler ProfilerConfig `yaml:"profiler"`
}

func Motd() string {
	return gconfig.Server.Motd
}

func Host() string {
	return gconfig.Server.Host
}

func Port() int {
	return gconfig.Server.Port
}

func IsOnline() bool {
	return gconfig.Server.Online
}

func Protocol() int {
	return gconfig.Server.Protocol
}

func LogLevel() LoggerLevel {
	switch strings.ToUpper(gconfig.Logger.Level) {
	case "OFF":
		return OFF
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return INFO
	}
}

type ServerConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Online   bool   `yaml:"online"`
	Motd     string `yaml:"motd"`
	Protocol int    `yaml:"protocol"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}

type ProfilerConfig struct {
	Port int `yaml:"port"`
}

func ProfilerPort() int {
	return gconfig.Profiler.Port
}

type LoggerLevel int

const (
	DEBUG LoggerLevel = iota
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

func Parse(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	yaml.Unmarshal([]byte(file), &gconfig)
}
