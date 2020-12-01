package main

import (
	"errors"
	"github.com/kattaris/errhand/pkg/errhand"
)

var log = errhand.New()
var logsPath = "D:/logs/testLogs.log"

func init() {
	log.CustomLogger(logsPath, "debug")
}

func main() {
	err := errors.New("test")

	for i := 0; i < 10; i++ {
		log.HandleSimpleErr(err)
		log.Infoln("Custom info")
		log.Debugln("Custom debug")
	}
}
