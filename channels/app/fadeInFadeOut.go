package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type InputPair [2]string
type SentencePair []InputPair

func generateInput(pairs SentencePair) <-chan InputPair {
	inChannel := make(chan InputPair, 2)
	go func() {
		defer close(inChannel)
		for _, pair := range pairs {
			inChannel <- pair
		}
	}()
	return inChannel
}

func decide(sentences InputPair) int {
	return len(strings.Split(sentences[0], " ")) + len(strings.Split(sentences[1], " "))
}

func FanOut(workerId int, producingChannel <-chan InputPair) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range producingChannel {
			fmt.Println(workerId, value[0], value[1])
			//fmt.Println(decide(value))
			out <- decide(value)
		}
	}()
	return out
}

func FanIn(workChannels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	mergeResult := make(chan int)
	outputChannel := func(o <-chan int) {
		defer wg.Done()
		for val := range o {
			mergeResult <- val
		}
	}
	wg.Add(len(workChannels))
	for _, ch := range workChannels {
		go outputChannel(ch)
	}
	go func() {
		wg.Wait()
		close(mergeResult)
	}()
	return mergeResult
}

func main() {
	var input SentencePair
	reader := bufio.NewScanner(os.Stdin)
GOPHER:
	fmt.Print("Enter first sentence of the pair:")
	reader.Scan()
	first := reader.Text()
	fmt.Print("Enter second sentence of the pair:")
	reader.Scan()
	second := reader.Text()
	input = append(input, [2]string{first, second})
	fmt.Println("Do you have more inputs? [y/n]")
	reader.Scan()
	var decide string = reader.Text()
	if strings.ToLower(string(decide)) == "y" {
		goto GOPHER
	}
	producingChannel := generateInput(input)
	// assign the channel inputs to workers -> Fan Out
	workerChan1 := FanOut(1, producingChannel)
	workerChan2 := FanOut(2, producingChannel)
	workerChan3 := FanOut(3, producingChannel)
	// Fan In to the output channel
	resultChan := FanIn(workerChan1, workerChan2, workerChan3)
	for rch := range resultChan {
		fmt.Println("Total length of the input sentences is", rch)
	}
}
