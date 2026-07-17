package main

import "fmt"

func AddCities(c []string) {
	c = append(c, "Ahmedabad")
	c = append(c, "Hyderabad")
}

func AddCitiesByRef(c *[]string) {
	*c = append(*c, "Bangalore")
	*c = append(*c, "Gurgaon")
	*c = append(*c, "Kochi")
	*c = append(*c, "Thiruvananthapuram")
}

func main() {
	var cities = []string{"Delhi", "Mumbai", "Kolkata", "Chennai"}
	fmt.Println("Slice example")
	fmt.Printf("Cities of India: %v\n", cities)
	fmt.Printf("Length of slice cities: %d\n", len(cities))
	fmt.Printf("Capacity of slice cities: %d\n", cap(cities))
	AddCities(cities)
	fmt.Printf("Length of slice cities: %d\n", len(cities))
	fmt.Printf("Capacity of slice cities: %d\n", cap(cities))
	AddCitiesByRef(&cities)
	fmt.Printf("Length of slice cities: %d\n", len(cities))
	fmt.Printf("Capacity of slice cities: %d\n", cap(cities))
	AddCitiesByRef(&cities)
	fmt.Printf("Length of slice cities: %d\n", len(cities))
	fmt.Printf("Capacity of slice cities: %d\n", cap(cities))
	AddCitiesByRef(&cities)
	fmt.Printf("Length of slice cities: %d\n", len(cities))
	fmt.Printf("Capacity of slice cities: %d\n", cap(cities))
	AddCitiesByRef(&cities)
	fmt.Printf("Length of slice cities: %d\n", len(cities))
	fmt.Printf("Capacity of slice cities: %d\n", cap(cities))
}
