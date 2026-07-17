package main

import "fmt"

func main() {
	var funcVar func(index uint64, value string) string
	//iterations := 10
	var valueArray = make(map[uint64]string)
	funcVar = func(index uint64, value string) string {
		fmt.Println(index)
		return value
	}
	valueArray[0] = "Girish eins"
	valueArray[1] = "Girish zwei"
	valueArray[2] = "Girish drei"
	valueArray[3] = "Girish vier"
	valueArray[4] = "Girish fünf"
	valueArray[5] = "Girish sechs"
	valueArray[6] = "Girish sieben"
	valueArray[7] = "Girish acht"
	valueArray[8] = "Girish nuen"
	valueArray[9] = "Girish zhen"
	for i, v := range valueArray {
		fmt.Println(funcVar(i, v))
	}
}
