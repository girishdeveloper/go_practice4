package main

import (
	"fmt"
	"log"
	"os"
)

/*
Please Build and execute
*/
func main() {
	dirEntries, err := os.ReadDir("/home/girish/Documents")
	raiseError(err)
	for _, v := range dirEntries {
		fmt.Println("Name:", v.Name())
		fmt.Println("Is Directory?", v.IsDir())
		fmt.Println("File mode:", v.Type().Perm())
		fileInfo, err := v.Info()
		raiseError(err)
		fmt.Println("File info:", fileInfo.Size(), fileInfo.Sys())
	}
	os.Exit(0)
}

func raiseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
