package main

import "fmt"

func main() {
	fmt.Println("Arrays declaration and usage.")
	var arr [3]uint8
	var roomnumbers [7]uint8
	arr1 := [4]int{12, 14, 15, 17}
	arr3 := [...]string{"Golang", "is", "great", "and", "Zig", "is", "better"}

	fmt.Println("Syntax:")
	fmt.Println("var arr [3]uint8")
	fmt.Println("var roomnumbers [7]uint8")
	fmt.Println("arr1 := [4]int{12,14,15,17}")
	fmt.Println("arr3 := [...]string{\"Golang\", \"is\", \"great\", \"and\", \"Zig\", \"is\", \"better\"}")

	for index, value := range arr1 {
		fmt.Println(index, value)
	}
	fmt.Printf("arr capacity= %d, array length= %d\n", cap(arr), len(arr))
	fmt.Printf("roomnumbers capacity= %d, array length= %d\n", cap(roomnumbers), len(roomnumbers))
	fmt.Printf("arr3 capacity= %d, array length= %d\n", cap(arr3), len(arr3))
	fmt.Println(arr3)
	arr3[2] = "work"
	fmt.Println(arr3)
	roomnumbers[0] = 3
	fmt.Printf("roomnumbers capacity= %d, array length= %d\n", cap(roomnumbers), len(roomnumbers))
	roomnumbers[1] = 4
	fmt.Printf("roomnumbers capacity= %d, array length= %d\n", cap(roomnumbers), len(roomnumbers))
	roomnumbers[2] = 5
	fmt.Printf("roomnumbers capacity= %d, array length= %d\n", cap(roomnumbers), len(roomnumbers))
}
