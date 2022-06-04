package errorHandler

import "log"

func LogErrorThenExit(location string, error error) {
	log.Fatalf("ERROR in %s: %v\n", location, error)
}
