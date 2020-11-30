package main

import (
	"errors"
	"github.com/kattaris/errhand/pkg/errhand"
)

var hndl = errhand.New("BOT_LOG", "debug")

func main() {
	err := errors.New("test")
	hndl.HandleSimpleErr(err)
}
