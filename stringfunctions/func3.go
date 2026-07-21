package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var str string = "   Foo   is    &  $@   Bar   "
	sliceString := strings.Fields(str)
	fmt.Printf("Fields in string %q are %q\n", str, sliceString)
	sliceString = strings.Fields("   Work  is  worship   ")
	fmt.Printf("Fields in string %q are %q\n", ("   Work  is  worship   "), sliceString)
	fry := func(c rune) bool {
		return (unicode.IsLetter(c) && !unicode.IsNumber(c))
	}
	fmt.Printf("FieldFunc %q is %q\n", str, strings.FieldsFunc(str, fry))
	// FieldsSeq
	text := "The quick brown fox"
	fmt.Println("Split sting text into fields")
	for word := range strings.FieldsSeq(text) {
		fmt.Printf("%q\n", word)
	}
	textWithSpaces := "     The    quick   brown   Fox     "
	fmt.Println("Split string with multiple spaces")
	for word := range strings.FieldsSeq(textWithSpaces) {
		fmt.Printf("%q\n", word)
	}
}
