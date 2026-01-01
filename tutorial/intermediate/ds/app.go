package main

func main() {
	var head *Node = nil;

	head = head.addFirst(12);
	head = head.addFirst(23);
	head = head.addFirst(34);

	head = head.addLast(45);
	head = head.addLast(56);
	head = head.addLast(67);

	head.display();
}