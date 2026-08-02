package main

import "fmt"

type Cart struct {
	Id          uint8
	ProductName string
	SKU         string
}

func (c Cart) getCart() Cart {
	return c
}

func (c *Cart) getCartByReference() Cart {
	return *c
}

func main() {
	fmt.Println("Method with Pointer receiver")
	c := &Cart{
		Id:          14,
		ProductName: "Noodles",
		SKU:         "BU21894784K",
	}
	fmt.Println("cart values are %v", c.getCart())
	d := *c
	d.Id = 15
	fmt.Println("cart values are %v", d.getCartByReference())
}
