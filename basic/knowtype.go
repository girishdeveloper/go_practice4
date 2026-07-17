package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("To know the type of a variable, type %T")
	fmt.Println("Syntax:")
	var variable1 uint64
	fmt.Println("var variable1 uint64")
	fmt.Printf("fmt.Printf(\"variable1 is of type % T\", variable1)\n")
	fmt.Printf("variable1 is of type %T\n", variable1)
	var str1 string
	str1 = "Work is worship"
	fmt.Println("str1 = \"Work is worship\"")
	fmt.Printf("reflect.TypeOf(str1)\n")
	fmt.Println(reflect.TypeOf(str1))
}
