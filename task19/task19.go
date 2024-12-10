package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "главрыба"
	reversed := reverseString(input)
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Перевернутая строка: %s\n", reversed)
}