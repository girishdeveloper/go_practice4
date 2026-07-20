package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("./files")
	if err != nil {
		panic(err)
	}
	fmt.Println("Removed a file")
}
