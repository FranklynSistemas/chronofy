package utils

import (
	"log"
)

func LogError(err error) {
	if err != nil {
		log.Printf("ERROR: %v\n", err)
	}
}
