package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(`Goroutines can communicate with each others through channels. 
	A channel can be seen as a pipeline of data between two goroutines. 
	This pipeline can only support a specific type.`)
	ch := make(chan int, 2)
	var received int

	fmt.Println("channel is", ch)
	ch <- 43
	received = <-ch
	ch <- 56
	ch <- 69
	received = <-ch
	fmt.Println("received value is", received)
	val, ok := <-ch
	if !ok {
		log.Println("channel value is not ok")
	}
	fmt.Println("value=", val)
	close(ch)
}
