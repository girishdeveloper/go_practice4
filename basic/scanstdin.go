package main

import "fmt"

func main() {
	var invar1 int64

	fmt.Println("To fetch input from the STDIN, use the below syntax.")
	fmt.Println("Note: specify the type identifier correctly.")
	fmt.Println("Syntax:")
	fmt.Println("var invar1 int64")
	fmt.Println("fmt.Print(\"Take input: \")")
	fmt.Println("fmt.Scanf(\"%d\", &invar1)")

	fmt.Print("Take input: ")
	fmt.Scanf("%d", &invar1)
	fmt.Printf("value %d scanned from standard input\n", invar1)
}
