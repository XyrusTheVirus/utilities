// @Title
// @Description
// @Author
// @Update

package stack

import (
	"fmt"
	"sync"

	datastructure "github.com/XyrusTheVirus/utilities/internal/dataStructure"
)

// Defines a stack data structure implemented by linked list
type Stack struct {
	datastructure.Common
}

func NewStack(maxCapacity uint) (*Stack, error) {
	stack := Stack{
		Common: datastructure.Common{
			Ds:          datastructure.NewLinkedList(datastructure.HEAD),
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

	s.Ds.(*datastructure.LinkedList).AddElement(val)
	s.Top = s.Ds.(*datastructure.LinkedList).GetHead()
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
	s.Top = s.Ds.(*datastructure.LinkedList).GetHead()
	s.NumOfElements--
	return temp.GetVal(), nil
}
