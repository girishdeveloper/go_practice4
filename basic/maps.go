package main

import "fmt"

func AddEntries(one map[string]uint8) {
	one["Girish"] = 13
	one["Koresh"] = 14
	fmt.Println("map1[\"Girish\"] = 13")
	fmt.Println("map1[\"Koresh\"] = 14")
}

func main() {
	var map1 = make(map[string]uint8)
	fmt.Println("var map1 = make(map[string]uint8)")
	fmt.Println("To add an element, follow the signature of the map variable")
	AddEntries(map1) // by default, pass-by-reference

	fmt.Println(map1)
	fmt.Println("To delete an item from a map,")
	delete(map1, "Koresh")
	fmt.Println("delete(map1,\"Koresh\")")
	fmt.Println(map1)
}
