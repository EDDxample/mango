package utils

import (
	"runtime"
	"strconv"
)

func GetPanicReportData() string {
	pc, _, line, ok := runtime.Caller(3)
	if ok {
		function := runtime.FuncForPC(pc)

		return function.Name() + ", line " + strconv.Itoa(line)
	} else {
		return "unknown, line unknown"
	}
}
