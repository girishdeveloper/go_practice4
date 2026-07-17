package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	x = 123
	y = 234
	z = x & y
	fmt.Printf("x=%d; y=%d; z=%d\n", x, y, z)
	z = x | y
	fmt.Printf("x=%d; y=%d; z=%d\n", x, y, z)
	z = x ^ y
	fmt.Printf("x=%d; y=%d; z=%d\n", x, y, z)
	z = x &^ y
	fmt.Printf("x=%d; y=%d; z=%d\n", x, y, z)
	z = x >> y
	fmt.Printf("x=%d; y=%d; z=%d\n", x, y, z)
	z = x << y
	fmt.Printf("x=%d; y=%d; z=%d\n", x, y, z)
}
