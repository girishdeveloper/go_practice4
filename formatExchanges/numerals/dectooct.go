package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	// Prompt the user
	fmt.Print("Enter an octal value (e.g., 755): ")
	fmt.Scan(&input)

	// Parse the string as base 8 into a 32-bit integer
	val, err := strconv.ParseInt(input, 8, 32)
	if err != nil {
		fmt.Printf("Invalid octal value: %v\n", err)
		return
	}

	// Assign to a variable
	var octalVariable uint32 = uint32(val)

	// Display the assigned value in decimal, standard octal, and Go 1.13+ octal formats
	fmt.Printf("Successfully assigned value!\n")
	fmt.Printf("Decimal: %d\n", octalVariable)
	fmt.Printf("Octal: %o\n", octalVariable)
	fmt.Printf("Go Literal: %#o\n", octalVariable) // Prefixes with 0o
}
