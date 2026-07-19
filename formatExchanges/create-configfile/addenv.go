package main

import (
	"fmt"
	"os"
)

func main() {
	myvar := os.Getenv("namespace")
	myvar2, found := os.LookupEnv("APP_PORT")
	if !found {
		fmt.Println("environment variable APP_PORT does not exist")
		myvar2 = "8081"
		os.Setenv("APP_PORT", myvar2)
	}
	fmt.Println("Namespace is ", myvar)
	fmt.Println("APP_PORT =", myvar2)
}
