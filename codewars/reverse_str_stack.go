package main

import (
	"fmt"
	"strings"
)

func main() {

	value := reverse("abcde") //edcba
	fmt.Println(value)
}

/*
func =>
i = string
o = reversed string

	A:
	- iterate through the string from the back
	- stack push

*/

type Stack struct {
	items []string
}

func (s *Stack) Push(char byte) {
	s.items = append(s.items, string(char))
}

func (s *Stack) Pop() {
	s.items = s.items[0 : len(s.items)-1]
}

func (s *Stack) Read() string {
	return string(s.items[len(s.items)-1])
}

func (s *Stack) ReadAll() string {
	return strings.Join(s.items, "")

}

func reverse(str string) string {
	myStack := Stack{}

	for idx := len(str) - 1; idx >= 0; idx-- {

		myStack.Push(str[idx])
	}
	return myStack.ReadAll()
}
