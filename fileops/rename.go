package main

import (
	"log"
	"os"
)

/*
Please Build and execute
*/
func main() {
	err := os.Rename(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}
