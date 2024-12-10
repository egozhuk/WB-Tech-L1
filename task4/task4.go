package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func producer(ctx context.Context, dataChan chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			num := rand.Intn(100)
			dataChan <- num
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker(ctx context.Context, id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: shutting down\n", id)
			return
		case num, ok := <-dataChan:
			if !ok {
				fmt.Printf("Worker %d: shutting down\n", id)
				return
			}
			fmt.Printf("Worker %d: received %d\n", id, num)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number_of_workers>")
		return
	}
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers")
		return
	}

	dataChan := make(chan int, 100)

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	go producer(ctx, dataChan)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataChan, &wg)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nShutting down...")

	cancel()
	wg.Wait()

	fmt.Println("All workers stopped. Exiting.")
}
