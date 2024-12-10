package main

import "fmt"

func intersect(set1, set2 []int) []int {
	elementMap := make(map[int]struct{})
	for _, elem := range set1 {
		elementMap[elem] = struct{}{}
	}

	intersection := []int{}
	for _, elem := range set2 {
		if _, exists := elementMap[elem]; exists {
			intersection = append(intersection, elem)
		}
	}

	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 7}
	set2 := []int{3, 4, 5, 6, 7}

	result := intersect(set1, set2)

	fmt.Printf("Пересечение множеств: %v\n", result)
}
