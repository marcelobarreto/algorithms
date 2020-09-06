package main

import "fmt"

type Queue struct {
	first *Node
	last  *Node
	size  uint
}

type Node struct {
	value int
	next  *Node
}

func (s *Queue) push(value int) uint {
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

func (s *Queue) pop() uint {
	if s.first == nil {
		return 0
	} else if s.size == 1 {
		s.first = nil
		s.last = nil
	} else {
		s.first = s.first.next
	}

	s.size--
	return s.size
}

func main() {
	queue := &Queue{}

	queue.push(1)
	fmt.Println("first:", queue.first.value, "last:", queue.last.value, "length:", queue.size)

	queue.push(2)
	fmt.Println("first:", queue.first.value, "last:", queue.last.value, "length:", queue.size)

	queue.push(3)
	fmt.Println("first:", queue.first.value, "last:", queue.last.value, "length:", queue.size)

	queue.pop()
	fmt.Println("first:", queue.first.value, "last:", queue.last.value, "length:", queue.size)

	queue.pop()
	fmt.Println("first:", queue.first.value, "last:", queue.last.value, "length:", queue.size)

	queue.pop()
	fmt.Println(queue)
}
