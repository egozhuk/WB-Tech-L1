package main

import (
	"fmt"
	"reflect"
)

func detectTypeWithReflect(v interface{}) {
	t := reflect.TypeOf(v)

	switch t.Kind() {
	case reflect.Int:
		fmt.Println("Тип: int")
	case reflect.String:
		fmt.Println("Тип: string")
	case reflect.Bool:
		fmt.Println("Тип: bool")
	case reflect.Chan:
		fmt.Println("Тип: chan")
	default:
		fmt.Println("Неизвестный или сложный тип")
	}
}

func main() {
	var intVar interface{} = 42
	var strVar interface{} = "Hello, Reflect!"
	var boolVar interface{} = true
	var intChan interface{} = make(chan int)

	detectTypeWithReflect(intVar)
	detectTypeWithReflect(strVar)
	detectTypeWithReflect(boolVar)
	detectTypeWithReflect(intChan)
	detectTypeWithReflect([]int{1, 2, 3})
}
