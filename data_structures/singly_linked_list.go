package main

import (
	"errors"
	"fmt"
)

// Node struct
type Node struct {
	val  int
	next *Node
}

// SinglyLinkedList struct
type SinglyLinkedList struct {
	length uint8
	head   *Node
	tail   *Node
}

// NewSinglyLinkedList is a SinglyLinkedList factory
func NewSinglyLinkedList() SinglyLinkedList {
	linkedList := SinglyLinkedList{}
	return linkedList
}

// SinglyLinkedList push add a new Node to the very end of a SLL
func (l *SinglyLinkedList) push(val int) *SinglyLinkedList {
	newNode := &Node{val: val}

	if l.head == nil {
		l.head = newNode
		l.tail = l.head
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}

	l.length++
	return l
}

// SinglyLinkedList pop remove the last Node of a SLL
func (l *SinglyLinkedList) pop() (*Node, error) {
	if l.head == nil {
		return nil, errors.New("pop error")
	}

	current := l.head
	newTail := current

	for current.next != nil {
		newTail = current
		current = current.next
	}

	l.tail = newTail
	l.tail.next = nil
	l.length--

	if l.length == 0 {
		l.head = nil
		l.tail = nil
	}

	return current, nil
}

// SinglyLinkedList shift remove first Node of a SLL
func (l *SinglyLinkedList) shift() (*Node, error) {
	if l.head == nil {
		return nil, errors.New("shift error")
	}

	currentHead := l.head
	l.head = currentHead.next
	l.length--

	if l.length == 0 {
		l.tail = nil
	}

	return currentHead, nil
}

// SinglyLinkedList unshift remove first Node of a SLL
func (l *SinglyLinkedList) unshift(val int) *Node {
	node := &Node{val: val}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}

	l.length++

	return node
}

// SinglyLinkedList unshift get Node at index of a SLL
func (l *SinglyLinkedList) get(index uint8) (*Node, error) {
	if index > l.length || l.length == 0 {
		return nil, errors.New("get could not be performed")
	}

	if index == 0 {
		return l.head, nil
	} else if index == l.length-1 {
		return l.tail, nil
	}

	var currentPosition uint8 = 0
	currentNode := l.head

	for currentPosition <= index {
		currentPosition++
		currentNode = currentNode.next
		if currentPosition >= index {
			break
		}
	}

	return currentNode, nil
}

// SinglyLinkedList set value at index of a SLL
func (l *SinglyLinkedList) set(index uint8, val int) bool {
	node, err := l.get(index)

	if err != nil {
		return false
	}

	node.val = val
	return true
}

// SinglyLinkedList insert a new Node between nodes of a SLL
func (l *SinglyLinkedList) insert(index uint8, val int) bool {
	if index > l.length {
		return false
	} else if index == l.length {
		l.push(val)
		return true
	} else if index == 0 {
		l.unshift(val)
		return true
	}

	node := &Node{val: val}
	prev, _ := l.get(index - 1)
	temp := prev.next
	prev.next = node
	node.next = temp
	l.length++

	return true
}

// SinglyLinkedList remove Node at index of a SLL
func (l *SinglyLinkedList) remove(index uint8) (*Node, error) {
	if index > l.length-1 {
		return nil, errors.New("error performing remove at index")
	} else if index == 0 {
		l.shift()
	} else if index == l.length-1 {
		l.pop()
	}

	prev, err := l.get(index - 1)

	if err != nil {
		return nil, err
	}

	node := prev.next
	prev.next = node.next
	l.length--

	return node, nil
}

// SinglyLinkedList reverse inverts a SLL
func (l *SinglyLinkedList) reverse() *SinglyLinkedList {
	node := l.head
	l.head = l.tail
	l.tail = node
	var next *Node
	var prev *Node

	for i := 1; uint8(i) < l.length; i++ {
		next = node.next
		node.next = prev
		prev = node
		node = next
	}

	return l
}

func log(statement string, l SinglyLinkedList) {
	fmt.Println(statement, "\n->", "Head:", l.head, "Tail:", l.tail, "Length:", l.length)
}

func main() {
	linkedList := NewSinglyLinkedList()

	linkedList.push(1)
	linkedList.push(2)
	linkedList.push(3)
	log("after_push 3x", linkedList)

	linkedList.pop()
	linkedList.pop()
	linkedList.pop()
	log("after_pop 3x", linkedList)

	linkedList.push(1)
	linkedList.push(2)
	linkedList.push(3)
	log("after_push 3x", linkedList)

	linkedList.shift()
	log("after_shift 1x", linkedList)

	linkedList.unshift(100)
	log("after_unshift 1x", linkedList)

	fmt.Println(linkedList.get(0))
	fmt.Println(linkedList.get(1))
	fmt.Println(linkedList.get(2))

	linkedList.set(0, 1)
	log("after_set at 0", linkedList)

	linkedList.set(0, 100)
	log("after_set at 0", linkedList)
	linkedList.set(1, 200)
	log("after_set at 1", linkedList)
	linkedList.set(2, 300)
	log("after_set at 2", linkedList)

	linkedList.insert(1, 150)
	log("after_insert at 1", linkedList)

	linkedList.remove(1)
	log("after_remove at 1", linkedList)

	fmt.Println(linkedList.get(1))

	linkedList.reverse()
	log("after_reverse", linkedList)
}
