package main

import (
	"encoding/xml"
	"fmt"
)

type Pet struct {
	Animal Cat `xml:"cat"`
}

type Cat struct {
	Name  string `xml:"name"`
	Breed string `xml:"breed"`
	Age   uint8  `xml:"age"`
}

func main() {
	pet := Pet{
		Animal: Cat{
			Name:  "sentri",
			Breed: "Mountain",
			Age:   25,
		},
	}
	fmt.Println("converting struct values to an xml")
	xml, err := xml.MarshalIndent(pet, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(xml))
}
