package main

import (
	"log"
	"os"
)

/*
BUILD and execute
*/
func main() {
	err := os.MkdirAll(os.Args[1], 0755)
	raiseError(err)
	os.Chmod(os.Args[1], 0777)
}

func raiseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
