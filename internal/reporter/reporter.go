package reporter

import (
	"log"
)

// HandleError is the centralized error reporting function.
func HandleError(err error) {
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
}
