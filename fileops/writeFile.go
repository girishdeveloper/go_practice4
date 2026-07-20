package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.WriteFile("./test.txt", []byte("Hello world!"), 0755)
	raiseError(err)

	var contents = make([]byte, 0)
	contents, err = os.ReadFile("./test.txt")
	raiseError(err)
	fmt.Println(string(contents))
}

func raiseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
