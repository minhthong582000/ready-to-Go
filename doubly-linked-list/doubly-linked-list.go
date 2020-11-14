package main

import (
	"errors"
	"fmt"
)

// Node in doubly linked list
type Node struct {
	data       interface{}
	next, prev *Node
}

// Node constructor
func node(data interface{}, next, prev *Node) *Node {
	var node Node = Node{data, next, prev}
	return &node
}

// DoublyLinkedList - Doubly Linked List structure
type DoublyLinkedList struct {
	head, tail *Node
	size       int
}

func (list *DoublyLinkedList) Clear() {
	var trav *Node = list.head
	for trav != nil {
		var next *Node = trav.next
		trav.prev, trav.next, trav.data = nil, nil, nil
		trav = next
	}

	list.head, list.tail, trav = nil, nil, nil
	list.size = 0
}

func (list DoublyLinkedList) GetSize() int {
	return list.size
}

func (list DoublyLinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *DoublyLinkedList) AddHead(data interface{}) {
	if list.IsEmpty() {
		newNode := node(data, nil, nil)
		list.head, list.tail = newNode, newNode
	} else {
		newNode := node(data, list.head, nil)
		list.head.prev = newNode
		list.head = list.head.prev
	}

	list.size++
}

func (list *DoublyLinkedList) AddLast(data interface{}) {
	if list.IsEmpty() {
		newNode := node(data, nil, nil)
		list.head, list.tail = newNode, newNode
	} else {
		newNode := node(data, nil, list.tail)
		list.tail.next = newNode
		list.tail = list.tail.next
	}

	list.size++
}

func (list DoublyLinkedList) Print() {
	var trav *Node = list.head

	fmt.Print("[ ")
	for trav != nil {
		fmt.Printf("%v ", trav.data)
		trav = trav.next
	}
	fmt.Print("]\n")
}

func (list DoublyLinkedList) PeekTail() interface{} {
	if list.IsEmpty() {
		return errors.New("list empty")
	}
	return list.tail.data
}

func (list DoublyLinkedList) PeekHead() interface{} {
	if list.IsEmpty() {
		return errors.New("list empty")
	}
	return list.head.data
}

func (list *DoublyLinkedList) RemoveHead() interface{} {
	if list.IsEmpty() {
		return errors.New("list empty")
	}

	var data interface{} = list.head.data
	list.head = list.head.next
	list.size--

	if list.IsEmpty() {
		list.tail = nil
	} else {
		list.head.prev = nil
	}

	return data
}

func (list *DoublyLinkedList) RemoveLast() interface{} {
	if list.IsEmpty() {
		return errors.New("list empty")
	}

	var data interface{} = list.tail.data
	list.tail = list.tail.prev
	list.size--

	if list.IsEmpty() {
		list.head = nil
	} else {
		list.tail.next = nil
	}

	return data
}

func (list *DoublyLinkedList) remove(node *Node) interface{} {
	// Handle if node is head or tail
	if node.next == nil { // It's tail
		return list.RemoveLast()
	}
	if node.prev == nil { // It's head
		return list.RemoveHead()
	}

	node.next.prev = node.prev
	node.prev.next = node.next

	data := node.data
	list.size--

	// Clean up data
	node.next, node.prev, node.data = nil, nil, nil

	return data
}

func (list *DoublyLinkedList) RemoveAt(index int) interface{} {
	if index < 0 || index >= list.size {
		return errors.New("invalid index")
	}

	var trav *Node

	if index < list.size/2 {
		trav = list.head
		for i := 0; index != i; i++ {
			trav = trav.next
		}
	} else {
		trav = list.tail
		for i := list.size - 1; index != i; i-- {
			trav = trav.prev
		}
	}

	return list.remove(trav)
}

func (list *DoublyLinkedList) Remove(obj interface{}) bool {
	var trav *Node = list.head

	for i := 0; i < list.size; i++ {
		if trav.data == obj {
			list.remove(trav)
			return true
		}
		trav = trav.next
	}

	return false
}

func (list *DoublyLinkedList) IndexOf(obj interface{}) int {
	var trav *Node = list.head

	for i := 0; i < list.size; i++ {
		if trav.data == obj {
			return i
		}
		trav = trav.next
	}

	return -1
}

func (list DoublyLinkedList) Contains(obj interface{}) bool {
	if list.IndexOf(obj) == -1 {
		return false
	}

	return true
}

func main() {
	var list DoublyLinkedList = DoublyLinkedList{nil, nil, 0}

	list.AddHead(11)
	list.AddLast(10)
	list.AddLast(1000)
	list.AddHead(9)

	list.Print()

	list.RemoveHead()
	list.RemoveHead()

	fmt.Println("-----------")

	fmt.Printf("removed: %v\n", list.RemoveLast())
	fmt.Printf("removed: %v\n", list.RemoveLast())
	fmt.Printf("removed: %v\n", list.RemoveLast())
	fmt.Printf("removed: %v\n", list.RemoveLast())

	list.Print()
	fmt.Println("-----------")

	list.AddLast(10)
	list.AddLast(20)
	list.AddHead(99)
	list.AddHead(nil)
	list.Print()
	fmt.Println("-----------")

	fmt.Printf("index of nil: %d\n", list.IndexOf(nil))
	fmt.Printf("index of 99: %d\n", list.IndexOf(99))
	fmt.Printf("List Contains nil: %t\n", list.Contains(nil))
	fmt.Printf("List Contains nigga: %t\n", list.Contains("nigga"))
	list.RemoveAt(4)
	list.Remove(nil)
	list.Print()
	fmt.Println("-----------")
}
