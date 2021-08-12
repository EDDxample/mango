package log

import "fmt"

var LOGGER = loggerClass{contextStack: []string{"start"}}

var colorReset string = "\033[0m"
var colorRed string = "\033[31m"
var colorGreen string = "\033[32m"
var colorYellow string = "\033[33m"
var colorBlue string = "\033[34m"
var colorPurple string = "\033[35m"
var colorCyan string = "\033[36m"
var colorWhite string = "\033[37m"

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
