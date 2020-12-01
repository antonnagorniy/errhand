package main

import (
	"errors"
	"github.com/kattaris/errhand/pkg/errhand"
)

var logger = errhand.Errhand{}
var logsPath = "D:/logs/testLogs.log"

func init() {
	logger.CustomLogger(logsPath, "debug")
}

func main() {
	err := errors.New("test")

	for i := 0; i < 10; i++ {
		logger.HandleSimpleErr(err)
		logger.Infoln("Custom info")
		logger.Debugln("Custom debug")
	}
}
