package main

import "fmt"

func main() {
	var aRune rune = 'G'

	fmt.Println("Rune is a character. \nIt can be converted into all but float and string datatype.")
	fmt.Println("Syntax:")
	fmt.Println("var aRune rune = 'G'")
	fmt.Println("fmt.Printf(\"Given rune is %c in character, %U in unicode, %d in integer\", aRune, aRune, aRune)")

	fmt.Printf("Given rune is %c in character, %U in unicode, %d in integer\n", aRune, aRune, aRune)
	strIn := "I like Golang"
	fmt.Println("Each character in a string can be treated as a rune.\nSyntax:")
	fmt.Printf("strIn := \"I like Golang\". Loop and print character by character.")
	for _, r := range strIn {
		fmt.Printf("rune: %c\n", r)
		fmt.Printf("rune is %c in character, %U in unicode, %d in integer\n", r, r, r)
	}
}
