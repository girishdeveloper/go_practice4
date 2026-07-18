package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Using Reader of strings package")

	var istr string
	fmt.Print("Enter a string: ")
	fmt.Scanf("%s", &istr)
	r := strings.NewReader(istr)
	fmt.Printf("Length of unread portion of the string %d\n", r.Len())
	var bt = []byte(istr)
	len, err := r.Read(bt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("length of read Byte %d\n", len)
	len, err = r.ReadAt(bt, 6)
	fmt.Printf("length of read Byte %d\n", len)
	byt, err := r.ReadByte()
	fmt.Printf("read Byte %v\n", byt)
	readrune, sizer, err := r.ReadRune()
	fmt.Printf("read Rune %v AND %d\n", readrune, sizer)
	r.Reset(istr)
	len, err = r.Read(bt)
	fmt.Printf("read %d\n", len)
	fmt.Printf("size %d\n", r.Size())
	r.Reset(istr)
}
