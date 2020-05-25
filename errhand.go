package errhand

import "log"

// HandleSimpleErr handles simple errors
func HandleSimpleErr(err error) {
	log.Fatalf("Error ocurred: %v\n", err)
}
