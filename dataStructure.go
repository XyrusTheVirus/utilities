// @Title
// @Description
// @Author
// @Update

package utilities

import (
	"math"
	"reflect"
	"sync"
)

type DataStructure interface {
	AddElement(val interface{}) interface{}
	DeleteElement(node *Node)
	Print()
}

type Common struct {
	Top           *Node
	NumOfElements uint
	Ds            DataStructure
	Mu            sync.Mutex
	MaxCapacity   uint
}

var maxMemorySpace uint = uint(math.Pow(2, 30))

// Returns the number of the elements
func (c *Common) NunOfElements() uint {
	return c.NumOfElements
}

// Returns if stack is empty or not
func (c *Common) IsEmpty() bool {
	return c.Top == nil
}

func (c *Common) Print() {
	c.Ds.Print()
}

func (c Common) IsMaximumMemoryExceeded() bool {
	return (uint(reflect.TypeOf(c).Size()) * c.MaxCapacity) > maxMemorySpace
}

func (c Common) IsMaxCapicityExcedded() bool {
	return (c.NumOfElements + 1) > c.MaxCapacity
}
