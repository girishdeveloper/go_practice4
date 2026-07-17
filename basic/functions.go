package main

import "fmt"

func funcWithNoParam() {
	fmt.Println("function without parameters")
}

func funcWith2Params(count int64, voice string) {
	fmt.Println("function with parameters")
	fmt.Println("parameter count", count)
	fmt.Println("parameter voice", voice)
}

func func3(index, length int64) int64 {
	var total int64 = index + length
	return total
}

func funcWithManyParams(par ...int64) (total int64) {
	total = 0
	for i := 0; i < len(par); i++ {
		total += par[i]
	}
	return
}

func funcWithMultipleReturn(a int64, b float64) (d float64, err string) {
	d = float64(a) + b
	err = ""
	if a < 5 {
		err = "This is not enough"
	}
	return d, err
}

func main() {
	fmt.Println("learn function declaration and functions.")
	funcWithNoParam()
	funcWith2Params(3, "Hello")
	fmt.Printf("total is %d\n", func3(3, 20))
	fmt.Printf("total of 3 values is %d\n", funcWithManyParams(12, 13, 154))
	fmt.Printf("total of 6 values is %d\n", funcWithManyParams(12, 13, 154, 16, 17, 20))
	c, err := funcWithMultipleReturn(3, 125.89)
	fmt.Printf("function returning two values are %3.2f and %s\n", c, err)
}
