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
	log.Println("received now 01")
	ch <- 69
	log.Println("received now 02")
	ch <- 73
	log.Println("received now 03")
	ch <- 87
	log.Println("received now 04")
	ch <- 77
	log.Println("received now 05")
	ch <- 98
	log.Println("received now 06")
	close(ch)
}

func chanOks(ch chan int) {
	fmt.Println("go routine starting", ch)
	/*val, ok := <-ch
	if !ok {
		log.Println("channel value is not ok")
	}*/
	for val := range ch {
		log.Println("value=", val)
	}
	fmt.Println("go routine end")
}
