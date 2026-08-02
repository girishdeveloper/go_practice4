package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func viewUnsignedInteger() {
	fmt.Println("Syntax:")
	fmt.Println("var variable1 uint32")
	fmt.Println("var variable2 uint64")
	fmt.Println("fmt.Scan(\"%d\", &variable2)")
	fmt.Println("fmt.Printf(\"%d\", variable2)")
}

func viewSignedInteger() {
	fmt.Println("Syntax:")
	fmt.Println("var variable1 int32")
	fmt.Println("var variable2 int64")
	fmt.Println("fmt.Scan(\"%d\", &variable2)")
	fmt.Println("fmt.Printf(\"%d\", variable2)")
}

func viewFloatingPoint() {
	fmt.Println("Syntax:")
	fmt.Println("var variable1 float32")
	fmt.Println("var variable2 float64")
	fmt.Println("fmt.Scan(\"%f\", &variable2)")
	fmt.Println("fmt.Printf(\"%f\", variable2)")
}

func viewHexaDecimal() {
	fmt.Println("Syntax:")
	fmt.Println("var variable1 = 0x124")
	fmt.Println("var variable2 = 0X124")
	fmt.Println("fmt.Scan(\"%x\", &variable2)")
	fmt.Println("fmt.Printf(\"%x\", variable2)")
}

func viewRune() {
	fmt.Println("Syntax:")
	fmt.Println("var variable1 rune = 'c'")
	fmt.Println("fmt.Scan(\"%v\", &variable1)")
	fmt.Println("fmt.Printf(\"Unicode code point %U, character %c, binary %b, hex %X, Decimal %d\", variable1, variable1, variable1, variable1, variable1)")
}

func viewString() {
	fmt.Println("Syntax:")
	fmt.Println("var variable1 string = \"I am Girish\"")
	fmt.Println("fmt.Sscan(\"%s\", &variable1)")
	fmt.Println("fmt.Sprintf(\"%s\", variable1)")
	fmt.Println("var str12 = `This is a string")
	fmt.Println("I typed a string`")
}

func main() {
	fmt.Println("This is the basic of go langauage")
	var idx int32 = 0
	var fvar float32 = 14.27
	fmt.Println("float variable =", fvar)
	for idx != -1 {
		fmt.Println("1. unsigned integer")
		fmt.Println("2. signed integer")
		fmt.Println("3. floating point")
		fmt.Println("4. Hexa-decimal")
		fmt.Println("5. Rune")
		fmt.Println("6. String")
		fmt.Print("Enter your selection: ")
		fmt.Scanf("%d", &idx)
		fmt.Printf("You chose %d\n", idx)
		switch idx {
		case 1:
			viewUnsignedInteger()
			break
		case 2:
			viewSignedInteger()
			break
		case 3:
			viewFloatingPoint()
			break
		case 4:
			viewHexaDecimal()
			break
		case 5:
			viewRune()
			break
		case 6:
			viewString()
			break
		} // end of switch
		s := fmt.Sprintln(strings.Repeat("*", 12))
		io.WriteString(os.Stdout, s)
	} // end of for loop
	fmt.Println("Ending this session")
}
