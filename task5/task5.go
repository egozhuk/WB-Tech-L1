package main

import (
	"fmt"
	"time"
)

func producer(dataChan chan int) {
	counter := 1
	for {
		dataChan <- counter
		counter++
		time.Sleep(500 * time.Millisecond)
	}
}

func consumer(dataChan chan int, doneChan chan struct{}) {
	for {
		select {
		case data := <-dataChan:
			fmt.Printf("Received: %d\n", data)
		case <-doneChan:
			fmt.Println("Consumer: shutting down")
			return
		}
	}
}

func main() {
	const duration = 5
	dataChan := make(chan int)
	doneChan := make(chan struct{})

	go producer(dataChan)

	go consumer(dataChan, doneChan)

	time.Sleep(time.Duration(duration) * time.Second)

	close(doneChan)
	fmt.Println("Main: shutting down")
}
