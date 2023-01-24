// @Title
// @Description
// @Author
// @Update

package queue

import (
	"fmt"
	"sync"

	datastructure "github.com/XyrusTheVirus/utilities/dataStructure"
)

// Defines a stack data structure implemented by linked list
type Queue struct {
	datastructure.Common
}

func NewQueue(maxCapacity uint) (*Queue, error) {
	queue := Queue{
		Common: datastructure.Common{
			Ds:          datastructure.NewLinkedList(datastructure.TAIL),
			Mu:          sync.Mutex{},
			MaxCapacity: maxCapacity,
		},
	}

	if queue.IsMaximumMemoryExceeded() {
		return nil, fmt.Errorf("Queue max capicity, exceeded Maximum memory space")
	}

	return &queue, nil
}

// Inserts item to the top of the stack
// Receives interface{} item to insert
func (q *Queue) Enqueue(val interface{}) error {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	if q.IsMaxCapicityExcedded() {
		return fmt.Errorf("Maximum capacity exceeded")
	}

	q.Ds.(*datastructure.LinkedList).AddElement(val)
	q.Top = q.Ds.(*datastructure.LinkedList).GetHead()
	q.NumOfElements++
	return nil
}

// Removes item from the top of the stack
func (q *Queue) Dequeue() (interface{}, error) {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	if q.IsEmpty() {
		return nil, fmt.Errorf("Queue is empty")
	}

	val := q.Top.GetVal()
	q.Ds.DeleteElement(q.Top)
	q.Top = q.Ds.(*datastructure.LinkedList).GetHead()
	q.NumOfElements--
	return val, nil
}
