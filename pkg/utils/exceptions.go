package utils

import "log"

// FatalErr is a function that logs a fatal error
func FatalErr(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

func LogErr(msg string, err error) {
	if err != nil {
		log.Default().Println(msg, err)
	}
}
