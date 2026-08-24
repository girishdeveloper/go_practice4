package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"slices"
	"time"

	"github.com/Rhymond/go-money"
)

type Pricing struct {
	UnitPrice float64
	Currency  string
}
type Product struct {
	Id       int
	ItemName string
	Price    Pricing
	Quantity float64
}
type Cart struct {
	Products []Product
	Total    float64
}

func (c *Cart) Init() {
	c.Products = make([]Product, 0)
	c.Total = 0.00
}

func (c *Cart) AddProduct(product Product) {
	c.Products = append(c.Products, product)
	c.Total = c.Total + product.Price.UnitPrice
}

func (c *Cart) List() {
	for i, v := range c.Products {
		fmt.Println(i+1, v)
	}
	fmt.Printf("total amount payable = %5.2f\n", c.Total)
}

func (c *Cart) Delete(index int) error {
	fmt.Println(index, reflect.TypeOf(c.Products[index]))
	if reflect.TypeOf(c.Products[index]) == nil {
		return errors.New("Unable to delete")
	}
	c.Total -= c.Products[index].Price.UnitPrice
	c.Products = slices.Delete(c.Products, index, index+1)
	fmt.Println("Deleted product", c.Products[index].ItemName)
	return nil
}

func main() {
	fmt.Println("This is an example of basic modular programming.")
	rand.NewSource(time.Now().UTC().UnixNano())
	kart := &Cart{}
	kart.Init()

	j := make([]int, 0)
	for i := 0; i < 5; i++ {
		j = append(j, rand.Intn(40))
		prod := Product{
			Id:       j[i],
			ItemName: "Maggi Noodles",
			Price: Pricing{
				UnitPrice: 20.49,
				Currency:  money.INR,
			},
			Quantity: 14,
		}
		kart.AddProduct(prod)
	}
	kart.List()
	errmsg := kart.Delete(rand.Intn(5))
	if errmsg != nil {
		fmt.Println(errmsg)
	}
	kart.List()
}
