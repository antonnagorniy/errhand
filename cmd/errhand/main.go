package main

import (
	"errors"
	"github.com/kattaris/errhand/v2/internal/handler"
)

var hndl = handler.New("BOT_LOG", "debug")

func main() {
	err := errors.New("test")

	hndl.HandleSimpleErr(err, "Got error!!!")
}
