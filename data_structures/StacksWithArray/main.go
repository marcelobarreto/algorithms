package main

import (
	"fmt"
)

type Stack struct {
	self []int
}

func (s *Stack) push(val int) int {
	s.self = append(s.self, val)
	return len(s.self)
}

func (s *Stack) pop() int {
	length := len(s.self)
	if length > 0 {
		s.self = append(s.self[:length-1])
		return len(s.self)
	}

	return 0
}

func main() {
	stack := &Stack{}
	fmt.Println(stack.push(1))
	fmt.Println(stack.self)
	fmt.Println(stack.push(2))
	fmt.Println(stack.self)
	fmt.Println(stack.pop())
	fmt.Println(stack.self)
	fmt.Println(stack.pop())
	fmt.Println(stack.self)
	fmt.Println(stack.pop())
	fmt.Println(stack.self)
}
