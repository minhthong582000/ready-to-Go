package main

import (
	"fmt"

	doublylinkedlist "github.com/minhthong582000/ready-to-Go/doubly-linked-list"
)

type Stack struct {
	List doublylinkedlist.DoublyLinkedList
}

func NewStack() Stack {
	var List doublylinkedlist.DoublyLinkedList = doublylinkedlist.NewDoublyLinkedList(nil, nil, 0)

	return Stack{
		List: List,
	}
}

func (s Stack) size() int {
	return s.List.GetSize()
}

func (s Stack) isEmpty() bool {
	return s.List.IsEmpty()
}

func (s *Stack) push(elem interface{}) {
	s.List.AddLast(elem)
}

func (s *Stack) pop() interface{} {
	return s.List.RemoveLast()
}

func (s *Stack) peek() interface{} {
	return s.List.PeekTail()
}

func (s Stack) print() {
	s.List.Print()
}

func isLeftBracket(brace string) bool {
	if brace == "{" || brace == "[" || brace == "(" {
		return true
	}

	return false
}

func getReverseBracket(bracket string) string {
	if bracket == "}" {
		return "{"
	}
	if bracket == "]" {
		return "["
	}
	if bracket == ")" {
		return "("
	}

	return ""
}

func BalanceBraces(braces string) bool {
	var stack Stack = NewStack()

	for _, char := range braces {
		if isLeftBracket(string(char)) {
			stack.push(string(char))
		} else if stack.isEmpty() || stack.pop() != getReverseBracket(string(char)) {
			return false
		}
	}

	return stack.List.IsEmpty()
}

func main() {
	// var stack Stack = NewStack()

	// fmt.Println(stack.size())
	// stack.push(1)
	// stack.push(1)
	// stack.push(1)

	// stack.print()

	// stack.pop()
	// stack.pop()
	// fmt.Println(stack.peek())
	// fmt.Println(stack.peek())
	// fmt.Println(stack.peek())
	// stack.print()

	fmt.Println(BalanceBraces("[][]()"))
}
