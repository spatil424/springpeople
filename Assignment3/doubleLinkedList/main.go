package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (dll *DoubleLinkedList) InsertAt(index int, value int) {
	if index < 0 || index > dll.size {
		panic(fmt.Sprintf("Error index %d out of bounds %d", index, dll.size))
	}

	newNode := &Node{data: value}
	if dll.size == 0 {
		dll.head = newNode
		dll.tail = newNode
	} else if index == 0 {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	} else if index == dll.size {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	} else {
		current := dll.head
		for i := 0; i < index; i++ {
			current = current.next
		}

		newNode.next = current
		newNode.prev = current.prev
		current.prev.next = newNode
		current.prev = newNode
	}
	dll.size++

}

func (dll *DoubleLinkedList) Sort() {
	if dll.head == nil || dll.head == dll.tail {
		return
	}

	for i := 0; i < dll.size-1; i++ {
		current := dll.head

		for j := 0; j < dll.size-1-i; j++ {
			if current.data > current.next.data {

				current.data, current.next.data = current.next.data, current.data
			}
			current = current.next
		}
	}
}

func (dll *DoubleLinkedList) Display() {
	if dll.head == nil {
		fmt.Println("List is empty.")
		return
	}
	current := dll.head
	fmt.Print("List (Head to Tail): ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		if current.next != nil {
			fmt.Print("<-> ")
		}
		current = current.next
	}
	fmt.Println()
}

func (dll *DoubleLinkedList) Reverse() {
	if dll.head == nil || dll.head == dll.tail {
		return
	}

	current := dll.head
	var tempNodeHolder *Node

	for current != nil {
		tempNodeHolder = current.prev
		current.prev = current.next
		current.next = tempNodeHolder
		current = current.prev
	}

	tempNodeHolder = dll.head
	dll.head = dll.tail
	dll.tail = tempNodeHolder
}

func recoverPanic() {
	r := recover()
	if r != nil {
		fmt.Println("Recovering from panic")
		fmt.Println("Error:  ", r)
	}
}

func main() {
	defer recoverPanic()
	fmt.Println("Double linked list demo")

	list := NewDoubleLinkedList()
	list.InsertAt(0, 10)
	list.InsertAt(1, 30)
	list.InsertAt(2, 5)
	list.InsertAt(3, 11)
	list.InsertAt(4, 81)
	fmt.Println("\nInitial list:")
	list.Display()

	fmt.Println("sorting the list")
	list.Sort()
	list.Display()

	fmt.Println("Reversing the list")
	list.Reverse()
	list.Display()
}
