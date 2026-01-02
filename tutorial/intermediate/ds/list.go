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

func (node *Node) delete() {
	node.Data = 0;
	node.Next = nil;
	node = nil;
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

func (node *Node) deleteCheck() *Node {
	if node == nil {
		return nil;
	} else if node.Next == nil {
		node = nil;
		return node;
	}
	return node;
}

func (node *Node) deleteFirst() *Node {
	curr := node.deleteCheck();

	if curr != nil {
		curr = node;
		node = node.Next;
		curr.delete();
	} else {
		fmt.Printf("E - empty list\n");
	}

	return node;
}

func (node *Node) deleteLast() *Node {
	curr := node.deleteCheck();

	if curr != nil {
		curr = node;
		var prev *Node;

		for curr.Next != nil {
			prev = curr;
			curr = curr.Next;
		}

		prev.Next = nil;
		curr.delete();
	} else {
		fmt.Printf("E - empty list\n");
	}

	return node;
}

//display helper function
func (node *Node) display() {
	curr := node;
	for curr != nil {
		fmt.Printf("%d ", curr.Data);
		curr = curr.Next;
	}
	fmt.Println();
}