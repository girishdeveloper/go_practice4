package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Paid float64 `json:"paid"`
}

func main() {
	fmt.Println("Exchange json data using a struct.")
	fmt.Println(strings.Repeat("*", 13))
	// create an object of type Product with values
	var p = Product{Id: 12, Name: "Girish"}
	fmt.Println("variable to json porting of values")
	bI, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.TypeOf(bI), string(bI))
	fmt.Println("Marshalling the json values (json encode)")
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
