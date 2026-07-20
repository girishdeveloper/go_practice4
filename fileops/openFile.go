package main

import (
	"log"
	"os"
)

func main() {
	fd, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_CREATE, 0755)
	raiseError(err)
	defer fd.Close()
	fd.WriteString("Hello Girish!")
}

func raiseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
