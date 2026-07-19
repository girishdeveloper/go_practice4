package main

import (
	"encoding/xml"
	"fmt"
)

type Animal struct {
	Horse `xml:"dog"`
}

type Horse struct {
	Name  string `xml:"name"`
	Breed string `xml:"breed"`
	Age   uint8  `xml:"age"`
}

func main() {
	prod := []byte(`<horse>
	<name>Alsafa</name>
	<breed>Alsatian</breed>
	<age>24</age>
	</horse>`)
	var pet = Animal{}
	fmt.Println("extract values from xml to a struct")
	err := xml.Unmarshal(prod, &pet)
	if err != nil {
		panic(err)
	}
	fmt.Println(pet)
}
