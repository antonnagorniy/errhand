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
	err := errors.New("test error")

	log.HandleError(err, false)
	log.Infoln(err)
	log.Debugln(err)
	log.Warnln(err)
	log.Fatalln(err)
}
