package main

import (
	"fmt"
	"os"
)

func main() {
	var head *Node
	var choice int

	for {
		fmt.Printf("1.Add 2.Rear Add 3.Delete 4.Rear Delete 5.Display 6.Exit: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			{
				var elem int
				fmt.Printf("Enter element: ")
				fmt.Scan(&elem)
				head = head.addFirst(elem)
			}
		case 2:
			{
				var elem int
				fmt.Printf("Enter element: ")
				fmt.Scan(&elem)
				head = head.addLast(elem)
			}
		case 3:
			{
				head = head.deleteFirst()
			}
		case 4:
			{
				head = head.deleteLast()
			}
		case 5:
			{
				head.display()
			}
		case 6:
			{
				os.Exit(0)
			}
		default:
			{
				fmt.Printf("Invalid choice\n")
			}
		}
	}
}
