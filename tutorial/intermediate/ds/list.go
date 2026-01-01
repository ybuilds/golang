package main

import "fmt"

type Node struct {
	Data int;
	Next *Node;
}

//constructor attached to type Node
func (node *Node) new(data int) *Node {
	return &Node{data, nil};
}

//add first function
func (node *Node) addFirst(data int) *Node{
	if node == nil {
		return node.new(data);
	}
	newNode := node.new(data);
	newNode.Next = node;
	return newNode;
}

//add last function
func (node *Node) addLast(data int) *Node {
	if(node == nil) {
		return node.new(data);
	}
	curr := node;
	for curr.Next != nil {
		curr = curr.Next;
	}
	curr.Next = node.new(data);
	return node;
}

//display helper function
func (node *Node) display() {
	curr := node;
	for curr != nil {
		fmt.Printf("%d ", curr.Data);
		curr = curr.Next;
	}
}