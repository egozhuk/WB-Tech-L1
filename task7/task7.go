package main

import (
	"fmt"
	"sync"
)

type ThreadSafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewThreadSafeMap() *ThreadSafeMap {
	return &ThreadSafeMap{
		data: make(map[string]int),
	}
}

func (m *ThreadSafeMap) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *ThreadSafeMap) Get(key string) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.data[key]
	return val, ok
}

func main() {
	safeMap := NewThreadSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key string, value int) {
			defer wg.Done()
			safeMap.Set(key, value)
			fmt.Printf("Записано: %s -> %d\n", key, value)
		}(fmt.Sprintf("key%d", i), i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		val, ok := safeMap.Get(key)
		if ok {
			fmt.Printf("Прочитано: %s -> %d\n", key, val)
		}
	}
}
