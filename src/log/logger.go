package log

import "fmt"

var LOGGER = loggerClass{contextStack: []string{"start"}}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type loggerClass struct {
	contextStack []string
}

func (logger loggerClass) Print(text string) {
	length := len(logger.contextStack) - 1
	fmt.Printf("[%s%s%s] %s\n", colorGreen, logger.contextStack[length], colorReset, text)
}

func (logger loggerClass) PrintError(err string) {
	length := len(logger.contextStack) - 1
	fmt.Printf("[%s%s%s] %s%s%s\n", colorRed, logger.contextStack[length], colorReset, colorYellow, err, colorReset)
}

func (logger *loggerClass) PushContext(context string) {
	logger.contextStack = append(logger.contextStack, context)
}

func (logger *loggerClass) PopContext() {
	length := len(logger.contextStack) - 1

	if length > 0 {
		logger.contextStack = logger.contextStack[:length]
	}
}
