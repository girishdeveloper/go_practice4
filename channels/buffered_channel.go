package main

import (
	"fmt"
	"log"
)

func main() {
	ch := make(chan int, 2)
	go chanOks(ch)
	log.Println("waiting for communication...")
	ch <- 56
	ch <- 69
	ch <- 73
	ch <- 87
	ch <- 77
	ch <- 98
	close(ch)
}

func chanOks(ch chan int) {
	fmt.Println("go routine starting", ch)
	/*val, ok := <-ch
	if !ok {
		log.Println("channel value is not ok")
	}*/
	for val := range ch {
		fmt.Println("value=", val)
	}
	fmt.Println("go routine end")
}
