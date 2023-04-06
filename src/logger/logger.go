package logger

import (
	"fmt"
	"mango/src/config"
	"os"
	"time"
)

func getTime() string {
	return time.Now().Format("15:04:05")
}

// Debug .
func Debug(args ...any) {
	if config.LogLevel() <= config.DEBUG {
		fmt.Printf("["+getTime()+"] [\x1B[94mDEBUG\x1B[0m] "+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)
	}
}

// Info .
func Info(args ...any) {
	if config.LogLevel() <= config.INFO {
		fmt.Printf("["+getTime()+"] [\x1B[92mINFO\x1B[0m] "+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)
	}
}

// Warn .
func Warn(args ...any) {
	if config.LogLevel() <= config.WARN {
		fmt.Printf("["+getTime()+"] [\x1B[33mWARN\x1B[0m] \x1B[33m"+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)
	}
}

// Error .
func Error(args ...any) {
	if config.LogLevel() <= config.ERROR {
		fmt.Printf("["+getTime()+"] [\x1B[91mERROR\x1B[0m] \x1B[91m"+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)
	}
}

// Fatal .
func Fatal(args ...any) {
	if config.LogLevel() <= config.FATAL {
		fmt.Printf("["+getTime()+"] [\x1B[91m\x1B[1mFATAL\x1B[0m] \x1B[91m\x1B[1m"+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)
	}

	os.Exit(1)
}
