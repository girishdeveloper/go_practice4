package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Usage of String functions")
	fmt.Println("strings package")
	var b strings.Builder
	fmt.Printf("Capacity of the builder %d\n", b.Cap())
	for i := 3; i > 0; i-- {
		fmt.Fprintf(&b, "%d,...", i)
	}
	b.WriteString("Intention")
	fmt.Fprintf(&b, "%d...", 12)
	fmt.Println(b.String())
	fmt.Printf("Capacity of the builder %d\n", b.Cap())
	b.Grow(40)
	fmt.Println("Grow the builder", b.Cap())
	fmt.Println("Length of the builder", b.Len())
	b.Reset()
	fmt.Println("Contents of the builder", b.String(), b.Cap(), b.Len())
	len, err := b.Write([]byte("This is Girish"))
	fmt.Println("length of builder", len, "value is", b.String(), "and error is", err)
	len, err = b.WriteRune('R')
	fmt.Println("adding a Rune to the builder", len, "value is", b.String(), "and error is", err)
}
