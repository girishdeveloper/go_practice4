package main

import (
	"encoding/json"
	"fmt"
)

type Storage interface {
	Read(name string, quantity float32) ([]byte, error)
	Write(content []byte)
}

type Material struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	Storage
}

func (m Material) Read(name string, quantity float32) ([]byte, error) {
	var content = Material{Name: name, Quantity: quantity}
	return json.MarshalIndent(content, "", " ")
}

func (m Material) Write(content []byte) {
	err := json.Unmarshal(content, &m)
	if err != nil {
		panic(err)
	}
	fmt.Println("data written", m)
}

func main() {
	var sto Material
	data, err := sto.Read("Washing Liquid", 345.03)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read the data")
	fmt.Println(string(data))
	writeData := []byte(`{"name": "Engine oil", "quantity": 549.09}`)
	sto.Write(writeData)
}
