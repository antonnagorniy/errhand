package main

import (
	"errors"
	"github.com/kattaris/errhand/v2"
)

func main() {
	err := errors.New("test")

	errhand.HandleSimpleErr(err, "Error in main:")

}
