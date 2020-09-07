package main

import "fmt"

/*
	BST (Binary Search Trees)
	Space  O(n)
	Insert O(log n)
	Remove O(log n)
	Search O(log n)
*/

// Node - BST's Node main struct
type Node struct {
	value int
	left  *Node
	right *Node
}

// BST - Binary Search Tree main struct
type BST struct {
	root *Node
}

// find - find a value from BST
func (b *BST) find(current *Node, value int) *Node {
	if current == nil {
		return nil
	} else if current.value == value {
		return current
	} else if value < current.value {
		return b.find(current.left, value)
	} else if value > current.value {
		return b.find(current.right, value)
	}

	return nil
}

// preInsert - Hook for searching pre-insert candidate
func (b *BST) findPreviousNode(current *Node, value int) *Node {
	if current == nil {
		current = b.root
		return b.findPreviousNode(current, value)
	} else if value < current.value && current.left != nil {
		return b.findPreviousNode(current.left, value)
	} else if value > current.value && current.right != nil {
		return b.findPreviousNode(current.right, value)
	}

	return current
}

// insert - adds a value to BST
func (b *BST) insert(value int) int {
	node := &Node{value: value}

	if b.root == nil {
		b.root = node
		return value
	}

	current := b.findPreviousNode(b.root, value)
	if value > current.value {
		current.right = node
	} else if value < current.value {
		current.left = node
	}

	return value
}

func main() {
	bst := &BST{}
	bst.insert(100)
	bst.insert(88)
	bst.insert(102)
	bst.insert(118)
	bst.insert(20)
	bst.insert(40)
	bst.insert(90)
	bst.insert(2)

	fmt.Println("trying to get 2 manually:", bst.root.left.left.left)
	fmt.Println("finding 91:", bst.find(bst.root, 91))
	fmt.Println("finding 90:", bst.find(bst.root, 90))
}
