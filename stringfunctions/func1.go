package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "work is worship"
	var str_clone string = strings.Clone(str)
	fmt.Println("original string content is", str)
	fmt.Println("cloned string is", str_clone)

	var str1 string = "compare this string"
	i := strings.Compare(str, str1)
	fmt.Println("Result: str > str1 i.e.", i)
	i = strings.Compare(str, str)
	fmt.Println("Result: str == str i.e.", i)
	i = strings.Compare(str1, str)
	fmt.Println("Result: str1 < str i.e.", i)

	if strings.Contains(str, "ship") == true {
		fmt.Println("str contains ship")
	}
	if strings.ContainsAny(str, "work") == true {
		fmt.Println("str contains work")
	}
	f := func(c rune) bool {
		return (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u')
	}
	if strings.ContainsFunc(str, f) == true {
		fmt.Println("str contains any of a,e,i,o and u. Found using function f")
	}
	if strings.ContainsRune(str, 'i') {
		fmt.Println("str contains rune i")
	}
	fmt.Println("How many occurance in str", strings.Count(str, "wor"))
	fmt.Println("Before and after each rune", strings.Count(str, ""), "!IMPORTANT")
}
