package errhand

import "log"

func HandleSimpleErr(err error) {
	log.Fatalf("Error ocurred: %v\n", err)
}
