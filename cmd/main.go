// @Title
// @Description
// @Author
// @Update

package main

import "test/utilities/stack"

func main() {
	stack := stack.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
}