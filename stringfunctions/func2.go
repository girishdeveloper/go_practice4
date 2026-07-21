package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "work is worship"
	before, after, found := strings.Cut(str, "is")
	fmt.Printf("Cut(%q, %q) = before=%q; after=%q; found=%v\n", str, "is", before, after, found)
	after, found = strings.CutPrefix(str, "wor")
	fmt.Printf("CutPrefix(%q, %q) = after=%q; found=%v\n", str, "wor", after, found)
	before, found = strings.CutSuffix(str, "ship")
	fmt.Printf("CutSuffix(%q, %q) = before=%q; found=%v\n", str, "ship", before, found)
	var str1 string = "Work Is Worship"
	if strings.EqualFold(str, str1) == true {
		fmt.Printf("%q is case-insensitively equal to %q\n", str, str1)
	}
	var str2 string = "stop SShipping"
	if strings.EqualFold(str, str2) == false {
		fmt.Printf("%q is not equal to %q\n", str, str2)
	}
}
