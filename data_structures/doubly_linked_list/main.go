package main

import (
	"fmt"
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	length uint
	head   *Node
	tail   *Node
}

func newDoublyLinkedList(values []int) DoublyLinkedList {
	d := DoublyLinkedList{}
	for i := 0; i < len(values); i++ {
		d.push(values[i])
	}

	return d
}

func (d *DoublyLinkedList) push(val int) *DoublyLinkedList {
	node := &Node{val: val}

	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		d.tail.next = node
		node.prev = d.tail
		d.tail = node
	}

	d.length++

	return d
}

func (d *DoublyLinkedList) pop() *Node {
	var node *Node

	if d.length == 1 {
		node = d.head
		d.head = nil
		d.tail = nil
	} else {
		node = d.tail
		d.tail = d.tail.prev
		d.tail.next = nil
	}

	d.length--
	return node
}

func (d *DoublyLinkedList) shift() *Node {
	node := d.head
	if d.length == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.head = node.next
		d.head.prev = nil
	}

	d.length--
	return node
}

func (d *DoublyLinkedList) unshift(val int) *DoublyLinkedList {
	node := &Node{val: val}
	if d.length == 0 {
		d.head = node
		d.tail = node
	} else {
		oldHead := d.head
		oldHead.prev = node
		node.next = oldHead
		d.head = node
	}

	d.length++
	return d
}

func (d *DoublyLinkedList) get(index uint8) *Node {
	if index == 0 {
		return d.head
	} else if index == uint8(d.length-1) {
		return d.tail
	}

	if index <= uint8(d.length/2) {
		current := d.head
		for i := 1; i <= int(index); i++ {
			current = current.next
		}
		return current
	}

	current := d.tail
	for i := int(d.length) - 1; i > int(index); i-- {
		current = current.prev
	}
	return current
}

func log(statement string, d DoublyLinkedList) {
	fmt.Println(statement, "\n->", "Head:", d.head, "Tail:", d.tail, "Length:", d.length)
}

func main() {
	linkedList := newDoublyLinkedList([]int{1, 2, 3, 4, 5})
	log("after_init", linkedList)

	linkedList.push(6)  // 1, 2, 3, 4, 5, 6
	linkedList.push(7)  // 1, 2, 3, 4, 5, 6, 7
	linkedList.push(8)  // 1, 2, 3, 4, 5, 6, 7, 8
	linkedList.push(9)  // 1, 2, 3, 4, 5, 6, 7, 8, 9
	linkedList.push(10) // 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	log("after_push", linkedList)

	linkedList.pop() // 1, 2, 3, 4, 5, 6, 7, 8, 9
	linkedList.pop() // 1, 2, 3, 4, 5, 6, 7, 8
	log("after_pop", linkedList)

	linkedList.shift() // 2, 3, 4, 5, 6, 7, 8
	log("after_shift", linkedList)

	linkedList.unshift(1) // 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	log("after_unshift", linkedList)

	fmt.Println("getting index: ", 0, linkedList.get(0)) // 1
	fmt.Println("getting index: ", 1, linkedList.get(1)) // 2
	fmt.Println("getting index: ", 2, linkedList.get(2)) // 3
	fmt.Println("getting index: ", 3, linkedList.get(3)) // 4
	fmt.Println("getting index: ", 4, linkedList.get(4)) // 5
	fmt.Println("getting index: ", 5, linkedList.get(5)) // 6
	fmt.Println("getting index: ", 6, linkedList.get(6)) // 7
	fmt.Println("getting index: ", 7, linkedList.get(7)) // 8
}
