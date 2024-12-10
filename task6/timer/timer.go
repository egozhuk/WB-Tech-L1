package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, stopChan <-chan struct{}, id int) {
	defer wg.Done()
	fmt.Printf("Воркер %d запущен.\n", id)
	for {
		select {
		case <-stopChan:
			fmt.Printf("Воркер %d завершает работу.\n", id)
			return
		default:
			fmt.Printf("Воркер %d выполняет задачу.\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var (
		wg         sync.WaitGroup
		numWorkers = 3
		duration   = 4 * time.Second
	)

	stopChan := make(chan struct{})

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, stopChan, i)
	}

	go func() {
		time.Sleep(duration)
		close(stopChan)
	}()

	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
