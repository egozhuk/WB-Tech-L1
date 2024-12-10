package main

import (
	"fmt"
	"time"
)

func worker(stopChan chan bool) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Worker: stopping")
			return
		default:
			fmt.Println("Worker: working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopChan := make(chan bool)

	go worker(stopChan)

	time.Sleep(2 * time.Second)
	stopChan <- true
	time.Sleep(1 * time.Second)
}
