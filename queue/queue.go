package main

import (
	"fmt"

	doublylinkedlist "github.com/minhthong582000/ready-to-Go/doubly-linked-list"
)

type Queue struct {
	list doublylinkedlist.DoublyLinkedList
}

func NewQueue() Queue {
	return Queue{
		list: doublylinkedlist.NewDoublyLinkedList(nil, nil, 0),
	}
}

func (q Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q Queue) Size() int {
	return q.list.GetSize()
}

func (q Queue) Peek() interface{} {
	return q.list.PeekHead()
}

func (q *Queue) Poll() interface{} {
	return q.list.RemoveHead()
}

func (q *Queue) Offer(elem interface{}) {
	q.list.AddLast(elem)
}

func (q Queue) Print() {
	q.list.Print()
}

func main() {
	q := NewQueue()

	q.Offer(1)
	q.Offer(2)
	q.Poll()
	q.Poll()
	q.Poll()
	fmt.Println(q.IsEmpty())
	q.Print()
	q.Offer(2)
	q.Offer("xxx")
	q.Offer("asdd")
	fmt.Println(q.Peek())
	fmt.Println(q.Poll())
	q.Print()
}
