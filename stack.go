// @Title
// @Description
// @Author
// @Update

package utilities

import (
	"fmt"
	"sync"
)

// Defines a stack data structure implemented by linked list
type Stack struct {
	Common
}

func NewStack(maxCapacity uint) (*Stack, error) {
	stack := Stack{
		Common: Common{
			Ds:          NewLinkedList(HEAD),
			Mu:          sync.Mutex{},
			MaxCapacity: maxCapacity,
		},
	}

	if stack.IsMaximumMemoryExceeded() {
		return nil, fmt.Errorf("Stack max capicity, exceeded Maximum memory space")
	}

	return &stack, nil
}

// Inserts item to the top of the stack
// Receives interface{} item to insert
func (s *Stack) Push(val interface{}) error {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	if s.IsMaxCapicityExcedded() {
		return fmt.Errorf("Maximum capacity exceeded")
	}

	s.Ds.(*LinkedList).AddElement(val)
	s.Top = s.Ds.(*LinkedList).GetHead()
	s.NumOfElements++
	return nil
}

// Removes item from the top of the stack
func (s *Stack) Pop() (interface{}, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	if s.IsEmpty() {
		return nil, fmt.Errorf("Stack is empty")
	}

	temp := s.Top
	s.Ds.DeleteElement(s.Top)
	s.Top = s.Ds.(*LinkedList).GetHead()
	s.NumOfElements--
	return temp.GetVal(), nil
}
