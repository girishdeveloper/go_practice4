package main

import (
	"encoding/json"
	"fmt"
)

type Pet struct {
	Dog `json:"dog"`
}

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint8  `json:"age"`
}

func main() {
	fmt.Println("pass json encoded string to fetch struct object with values")
	strjsonencoded := []byte(`{"dog": {"name": "Ponty", "breed": "Pitbull", "age": 42}}`)
	var pet = Pet{}
	fmt.Println("Unmarshal the value string (json decode)")
	err := json.Unmarshal(strjsonencoded, &pet)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Values are Name is %s, Breed is %s, Age is %d\n", pet.Dog.Name, pet.Dog.Breed, pet.Dog.Age)
}
