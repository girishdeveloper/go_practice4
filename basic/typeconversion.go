package main

import "fmt"

func main() {
	fmt.Println("To convert the value from one type to another, simply specify the %type-identifier in the printing statement")
	var var1 uint32 = 928
	var var2 uint64 = 238
	fmt.Println("Syntax:")
	fmt.Println("var var1 uint32 = 928")
	fmt.Println("var var2 uint64 = 238")
	fmt.Println("fmt.Printf(\"variable var1 has value %d in 32-bit integer, %f in float, %b in binary, %s in string, %U in unicode, %x in hexa-decimal\", var1, var1, var1, var1, var1, var1)")
	fmt.Println("fmt.Printf(\"variable var2 has value %d in 64-bit integer, %f in float, %b in binary, %s in string, %U in unicode, %x in hexa-decimal\", var2, var2, var2, var2, var2, var2)")

	fmt.Printf("variable var1 has value %d in 32-bit integer, %f in float, %b in binary, %s in string, %U in unicode, %x in hexa-decimal, %q\n", var1, var1, var1, var1, var1, var1, var1)
	fmt.Printf("variable var2 has value %d in 64-bit integer, %f in float, %b in binary, %s in string, %U in unicode, %x in hexa-decimal, %q\n", var2, var2, var2, var2, var2, var2, var2)
}
