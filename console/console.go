package console

import (
	"log"
	"os"
)

const logWithTimestampFlag = 3

var Logger = *log.New(os.Stdout, "[INFO] ", logWithTimestampFlag)
var LoggerWarning = *log.New(os.Stderr, "[WARN] ", logWithTimestampFlag)

func IfErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func IfErrorPrintWarningNoPanic(err error) {
	if err != nil {
		LoggerWarning.Println(err.Error())
	}
}
