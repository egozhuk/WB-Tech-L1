package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}

	for i, val := range arr {
		if i == len(arr)/2 {
			continue
		}
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)

	sortedArr := quickSort(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}
