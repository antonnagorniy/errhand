package main

import (
	"errors"
	"github.com/kattaris/errhand"
)

func main() {
	err := errors.New("test")

	errhand.HandleSimpleErr(err, "Error in main: ")
}
