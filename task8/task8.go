package main

import (
	"errors"
	"fmt"
)

func updateBit(n int64, i int, action string) (int64, error) {
	if i < 0 || i > 63 {
		return n, errors.New("индекс бита должен быть в диапазоне от 0 до 63")
	}

	switch action {
	case "set":
		return n | (1 << i), nil
	case "clear":
		return n & ^(1 << i), nil
	default:
		return n, errors.New(`недопустимое действие: используйте "set" или "clear"`)
	}
}

func main() {
	var n int64 = 42
	bitIndex := 3

	fmt.Printf("Число до изменений: %d (%b)\n", n, n)

	if result, err := updateBit(n, bitIndex, "set"); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		n = result
		fmt.Printf("После установки %d-го бита в 1: %d (%b)\n", bitIndex, n, n)
	}

	if result, err := updateBit(n, bitIndex, "clear"); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		n = result
		fmt.Printf("После установки %d-го бита в 0: %d (%b)\n", bitIndex, n, n)
	}
}
