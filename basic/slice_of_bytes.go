package main

import "fmt"

func main() {
	sliceBytes := make([]byte, 0)
	sliceBytes = append(sliceBytes, 4)
	sliceBytes = append(sliceBytes, 6)
	sliceBytes = append(sliceBytes, 34)
	fmt.Printf("type is %T and values are %v\n", sliceBytes, sliceBytes)
}
