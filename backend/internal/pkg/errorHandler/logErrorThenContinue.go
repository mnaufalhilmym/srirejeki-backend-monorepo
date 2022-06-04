package errorHandler

import "log"

func LogErrorThenContinue(location string, error interface{}) {
	log.Printf("ERROR in %s: %v\n", location, error)
}
