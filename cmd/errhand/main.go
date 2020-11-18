package main

import (
	"errors"
	"github.com/kattaris/errhand/pkg"
)

var hndl = pkg.New("BOT_LOG", "debug")

func main() {
	err := errors.New("test")
	hndl.HandleSimpleErr(err, "Got error: ")
}
