package main

import (
	"bufio"
	"fmt"
	"os"
)

type Country struct {
	Name        string
	CapitalCity string
}

type GoverningBody struct {
	Type  string
	Place string
}

func main() {
	fmt.Println("Play with struct")
	var bioScanner = bufio.NewScanner(os.Stdin)
	japan := Country{
		"Japan",
		"Tokyo",
	}
	fmt.Println("sample", japan)
	nation := GoverningBody{}
	fmt.Print("Governing body type: ")
	bioScanner.Scan()
	nation.Type = bioScanner.Text()
	fmt.Print("Place: ")
	bioScanner.Scan()
	nation.Place = bioScanner.Text()
	if err := bioScanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Result from the scanning is: ", nation)
}
