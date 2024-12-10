package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, running *bool, mu *sync.Mutex) {
	defer wg.Done()
	fmt.Printf("Воркер %d запущен.\n", id)
	for {
		mu.Lock()
		if !*running {
			mu.Unlock()
			fmt.Printf("Воркер %d завершает работу.\n", id)
			return
		}
		mu.Unlock()

		fmt.Printf("Воркер %d выполняет задачу.\n", id)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var (
		numWorkers = 3
		running    = true
		mu         sync.Mutex
		wg         sync.WaitGroup
	)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i, &running, &mu)
	}

	time.Sleep(5 * time.Second)

	mu.Lock()
	running = false
	mu.Unlock()

	wg.Wait()
	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}
