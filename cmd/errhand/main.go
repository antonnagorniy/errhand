package main

import (
	"errors"
	"github.com/kattaris/errhand/pkg"
)

var log *errhand.Errhand

func init() {
	log = errhand.GetLogger()
}

func main() {
	err := errors.New("test")

	log.HandleError(err, false)
	log.Infoln("Custom info")
	log.Debugln("Custom debug")
	log.Warnln("Custom warn")
	log.Fatalln("Custom fatal")
}
