package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Speak() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

type Action struct {
	Human
	ActionType string
}

func (a Action) PerformAction() {
	fmt.Printf("%s is performing action: %s\n", a.Name, a.ActionType)
	a.Speak()
}

func main() {
	person := Action{
		Human:      Human{Name: "John", Age: 30},
		ActionType: "running",
	}

	person.PerformAction()
}
