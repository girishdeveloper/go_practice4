package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("string Replacer functions")
	var str string = "Girish-Madhavan"
	repl := strings.NewReplacer("G", "g", "-", " ", "M", "m")
	str = repl.Replace(str)
	fmt.Printf("replacer %s\n", str)
}
