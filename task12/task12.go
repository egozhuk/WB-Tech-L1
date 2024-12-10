package main

import "fmt"

type StringSet struct {
	data map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{
		data: make(map[string]struct{}),
	}
}

func (s *StringSet) Add(value string) {
	s.data[value] = struct{}{}
}

func (s *StringSet) Remove(value string) {
	delete(s.data, value)
}

func (s *StringSet) Contains(value string) bool {
	_, exists := s.data[value]
	return exists
}

func (s *StringSet) Elements() []string {
	elements := make([]string, 0, len(s.data))
	for key := range s.data {
		elements = append(elements, key)
	}
	return elements
}

func (s *StringSet) Size() int {
	return len(s.data)
}

func main() {
	set := NewStringSet()

	set.Add("cat")
	set.Add("dog")
	set.Add("tree")
	set.Add("cat")

	fmt.Println("Множество содержит:")
	for _, elem := range set.Elements() {
		fmt.Println(elem)
	}

	fmt.Println("Содержит 'dog'? ", set.Contains("dog"))
	fmt.Println("Содержит 'bird'? ", set.Contains("bird"))

	set.Remove("dog")
	fmt.Println("После удаления 'dog':", set.Elements())

	fmt.Println("Размер множества:", set.Size())
}
