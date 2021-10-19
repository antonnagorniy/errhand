package main

import (
	"errors"
	"fmt"
	"github.com/kattaris/errhand"
	"os"
)

func init() {
	userName := os.Getenv("USER")
	logs := fmt.Sprintf("/home/%s/Logs/errhand/log.log", userName)
	errhand.Init(logs)
}

func main() {
	log := errhand.GetLogger()
	err := errors.New("test")

	log.HandleError(err, false)
	log.Infoln("Custom info")
	log.Debugln("Custom debug")
	log.Warnln("Custom warn")
	log.Fatalln("Custom fatal")
}
