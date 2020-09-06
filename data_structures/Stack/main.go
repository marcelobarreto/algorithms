package main

import "fmt"

type Stack struct {
	first *Node
	last  *Node
	size  uint
}

type Node struct {
	value int
	next  *Node
}

func (s *Stack) push(value int) uint {
	node := &Node{value: value}

	if s.size == 0 {
		s.last = node
		s.first = node
	} else if s.size == 1 {
		s.first.next = node
		s.last = node
	} else {
		s.last.next = node
		s.last = node
	}

	s.size++
	return s.size
}

func (s *Stack) pop() uint {
	if s.first == nil {
		return 0
	} else if s.size == 1 {
		s.first = nil
		s.last = nil
	} else {
		s.last = nil
		current := s.first

		for i := 0; i < int(s.size-2); i++ {
			current = current.next
		}
		current.next = nil
		s.last = current
	}

	s.size--
	return s.size
}

func main() {
	stack := &Stack{}

	stack.push(1)
	fmt.Println(stack)
	fmt.Println(stack.first.value, stack.last.value, stack.size)

	stack.push(2)
	fmt.Println(stack)
	fmt.Println(stack.first.value, stack.last.value, stack.size)

	stack.push(3)
	fmt.Println(stack)
	fmt.Println(stack.first.value, stack.last.value, stack.size)

	stack.pop()
	fmt.Println(stack)
	fmt.Println(stack.first.value, stack.last.value, stack.size)

	stack.pop()
	fmt.Println(stack)
	fmt.Println(stack.first.value, stack.last.value, stack.size)

	stack.pop()
	fmt.Println(stack)
}
