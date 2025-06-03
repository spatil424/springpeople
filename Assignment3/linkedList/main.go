package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

func (ll *LinkedList) Add(value int) {
	newNode := &Node{
		data: value,
		next: nil,
	}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
	ll.size++
}

func (ll *LinkedList) PrintList() {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}

	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}

func (ll *LinkedList) SortList() {
	if ll.head == nil || ll.head.next == nil {
		fmt.Println("Empty list or list with one element")
		return
	}

	for i := 0; i < ll.size-1; i++ {
		current := ll.head

		for j := 0; j < ll.size-1; j++ {
			if current != nil && current.next != nil {
				if current.data > current.next.data {
					current.data, current.next.data = current.next.data, current.data
				}
				current = current.next
			} else {
				break
			}
		}
	}

}

func (ll *LinkedList) Reverse() {
	if ll.head == nil || ll.head.next == nil {
		fmt.Println("Empty list or list with one element")
		return
	}

	var prev *Node = nil
	current := ll.head
	var nextNode *Node = nil

	for current != nil {
		nextNode = current.next
		current.next = prev
		prev = current
		current = nextNode
	}
	ll.head = prev

}

func main() {
	fmt.Println("Single Linked list implementation")

	list := NewLinkedList()

	list.Add(11)
	list.Add(21)
	list.Add(9)
	list.Add(77)

	fmt.Println("Before sorting")
	list.PrintList()

	fmt.Println("\nAfter sorting")
	list.SortList()
	list.PrintList()

	fmt.Println("\nAfter reversing")
	list.Reverse()
	list.PrintList()

}
