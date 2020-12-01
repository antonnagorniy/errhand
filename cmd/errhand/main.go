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

	log.HandleSimpleErr(err)
	log.Infoln("Custom info")
	log.Debugln("Custom debug")
	log.Warnln("Custom warn")
	log.Fatalln("Custom warn")
}
