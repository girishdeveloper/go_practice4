package main

import (
	"fmt"
)

func main() {
	fmt.Println("There are two types of strings.")
	fmt.Println("01.\t Raw strings -> strings between backquotes")
	fmt.Println("\t\tForbidden characters: backticks/backquotes")
	fmt.Println("\t\tDiscarded characters: carriage return \\r")
	fmt.Println("02. interpreted strings -> strings between double quotes")
	fmt.Println("\t\tForbidden characters: backquotes and carriage return \\n")
}
