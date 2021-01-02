package utils

import "log"

func Error(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
