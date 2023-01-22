// @Title
// @Description
// @Author
// @Update

package stack

import (
	"fmt"
	"sync"
	datastructure "test/utilities/internal/dataStructure"
)

// Defines a stack data structure implemented by linked list
type Stack struct {
	top           *datastructure.Node // Denotes the top of the stack
	numOfElements int
	ds            datastructure.DataStructure
	mu            sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		ds: datastructure.NewLinkedList("Head"),
		mu: sync.Mutex{},
	}
}

// Inserts item to the top of the stack
// Receives interface{} item to insert
func (s *Stack) Push(v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.top = s.ds.AddElement(v).(*datastructure.Node)
	s.numOfElements++
	return
}

// Removes item from the top of the stack
// Receives the head of the linked list
func (s *Stack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	temp := s.Top()
	s.ds.DeleteElement(s.Top())
	s.top = s.ds.(datastructure.LinkedList).GetHead()
	s.numOfElements--
	return temp.(*datastructure.Node).GetVal()
}

// Returns the top of the stack
func (s *Stack) Top() *datastructure.Node {
	return s.top
}

// Returns the number of the elements
func (s *Stack) NunOfElements() int {
	return s.numOfElements
}

// Returns if stack is empty or not
func (s *Stack) isEmpty() bool {
	return s.top == nil
}

// Prints the stack
// Receives the head of the linked list
func (s *Stack) Print() {
	temp := s.Top()
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}
