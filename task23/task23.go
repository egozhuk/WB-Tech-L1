package main

import "fmt"

func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Индекс вне диапазона")
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный слайс:", slice)

	index := 2
	result := removeElement(slice, index)

	fmt.Printf("После удаления %d-го элемента: %v\n", index, result)
}
