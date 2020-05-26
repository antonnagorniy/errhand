package errhand

import "log"

// HandleSimpleErr handles simple errors
func HandleSimpleErr(err error, message string) {
	if err != nil {
		log.Fatalln(message, err)
	}
}
