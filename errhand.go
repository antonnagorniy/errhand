package errhand

import "log"

// HandleSimpleErr handles simple errors
func HandleSimpleErr(err error, message string) {
	log.Fatalln(message, err)
}
