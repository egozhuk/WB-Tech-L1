package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Начало...")
	Sleep(2 * time.Second)
	fmt.Println("Прошло 2 секунды.")
}
