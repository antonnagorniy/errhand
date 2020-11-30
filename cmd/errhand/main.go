package main

import (
	"errors"
	"github.com/kattaris/errhand/pkg/errhand"
)

var stringPathHandler *errhand.Handler
var logsPath = "D:/logs/testLogs.log"

func init() {
	stringPathHandler = errhand.New(logsPath, "debug")
}

func main() {
	err := errors.New("test")
	stringPathHandler.HandleSimpleErr(err)
}
