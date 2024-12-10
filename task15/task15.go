package main

import "fmt"

/*
Строки в Go являются неизменяемыми и хранятся как структура, содержащая указатель на данные и длину.
Когда выполняется операция v[:100], создается новая строка, которая ссылается на те же данные, что и v.
Если v — очень большая строка, то её данные остаются в памяти, пока срез v[:100] существует, даже если
используется только небольшой кусок строки.
*/

func createHugeString(size int) string {
	result := make([]rune, size)
	for i := 0; i < size; i++ {
		result[i] = 'a'
	}
	return string(result)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	buf := make([]byte, 100)
	copy(buf, v[:100])
	justString = string(buf)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
