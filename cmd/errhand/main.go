package main

import (
	"errors"
	errhand2 "github.com/kattaris/errhand"
	"os"
)

var log = errhand2.New()
var logsPath string

func init() {
	logsPath = os.Getenv("TEST_LOGS") + "errhand/log.log"
	log.Infoln(logsPath)
	log.CustomLogger(logsPath, "debug")
}

func main() {
	err := errors.New("test")

	log.HandleError(err, false)
	log.Infoln("Custom info")
	log.Debugln("Custom debug")
	log.Warnln("Custom warn")
	log.Fatalln("Custom fatal")
}
