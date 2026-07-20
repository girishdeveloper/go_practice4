package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Get the numeric group id of the caller::", os.Getgid())
	fmt.Println("Get the numeric user id of the caller::", os.Getuid())
	fmt.Println("Get the process id of the caller (this process):", os.Getpid())
	fmt.Println("Get the process id of the parent process:", os.Getppid())
	pwd, err := os.Getwd()
	raiseError(err)
	fmt.Println("current working directory:", pwd)
	hostname, err := os.Hostname()
	raiseError(err)
	fmt.Println("Hostname of the current machine:", hostname)
}

func raiseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
