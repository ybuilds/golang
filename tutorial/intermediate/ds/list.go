package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

// constructor attached to type Node
func (node *Node) new(data int) *Node {
	return &Node{data, nil}
}

func (node *Node) delete() {
	node = nil
}

// add first function
func (node *Node) addFirst(data int) *Node {
	if node == nil {
		return node.new(data)
	}

	newNode := node.new(data)
	newNode.Next = node
	fmt.Printf("Added %d\n", newNode.Data)

	return newNode
}

// add last function
func (node *Node) addLast(data int) *Node {
	if node == nil {
		return node.new(data)
	}

	curr := node
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node.new(data)
	fmt.Printf("Added %d\n", curr.Next.Data)

	return node
}

func (node *Node) deleteFirst() *Node {
	if node == nil {
		fmt.Printf("List is empty\n")
		return nil
	} else if node.Next == nil {
		fmt.Printf("Deleted %d\n", node.Data)
		node = nil
		return node
	}

	curr := node
	node = node.Next
	fmt.Printf("Deleted %d\n", curr.Data)
	curr.delete()

	return node
}

func (node *Node) deleteLast() *Node {
	if node == nil {
		fmt.Printf("List is empty\n")
		return nil
	} else if node.Next == nil {
		fmt.Printf("Deleted %d\n", node.Data)
		node = nil
		return node
	}

	curr := node
	var prev *Node

	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	prev.Next = nil
	curr.delete()
	fmt.Printf("Deleted %d\n", curr.Data)

	return node
}

// display helper function
func (node *Node) display() {
	if node == nil {
		fmt.Printf("List is empty\n")
		return
	}
	fmt.Printf("List: ")
	curr := node
	for curr != nil {
		fmt.Printf("%d ", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}
