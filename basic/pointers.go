package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Id        uint64
	Name      string
	UnitPrice float64
	Currency  string
	Stock     uint64
}

func main() {
	var ptr *Product
	var value Product
	fmt.Println("Point-ers")
	ptr = &value
	value.Id = 153
	value.Name = "Girish"
	value.UnitPrice = 140.54
	value.Currency = "INR"
	value.Stock = 1743
	fmt.Println(strings.Repeat("*", 13))
	fmt.Printf("Product details are: %v; %v\n", *ptr, ptr.Id)
	value.Name = "Bajra"
	fmt.Printf("Product details are: %v; %v\n", *ptr, ptr.Name)
}
