package main

import (
	"fmt"
	"sync"
)

func multiplyByTwo(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := range input {
		output <- x * 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	inputChan := make(chan int)
	outputChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go multiplyByTwo(inputChan, outputChan, &wg)

	go func() {
		for result := range outputChan {
			fmt.Println(result)
		}
	}()

	go func() {
		for _, num := range numbers {
			inputChan <- num
		}
		close(inputChan)
	}()

	wg.Wait()
	close(outputChan)
}
