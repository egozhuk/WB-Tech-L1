package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	charMap := make(map[rune]bool)

	for _, char := range strings.ToLower(s) {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}
	return true
}

func main() {
	tests := []string{
		"abcd",      // true
		"abCdefAaf", // false
		"aabcd",     // false
		"unique",    // false
		"Uniqe",     // true
	}

	for _, test := range tests {
		fmt.Printf("Строка \"%s\" уникальна? %v\n", test, isUnique(test))
	}
}
