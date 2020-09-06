package main

import (
	"fmt"
)

type Queue struct {
	self []int
}

func (s *Queue) push(val int) int {
	s.self = append(s.self, val)
	return len(s.self)
}

func (s *Queue) pop() int {
	length := len(s.self)
	if length > 0 {
		s.self = append(s.self[1:])
		return len(s.self)
	}

	return 0
}

func main() {
	queue := &Queue{}
	fmt.Println(queue.push(1))
	fmt.Println(queue.self)
	fmt.Println(queue.push(2))
	fmt.Println(queue.self)
	fmt.Println(queue.pop())
	fmt.Println(queue.self)
	fmt.Println(queue.pop())
	fmt.Println(queue.self)
	fmt.Println(queue.pop())
	fmt.Println(queue.self)
}
