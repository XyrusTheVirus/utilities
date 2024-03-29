// @Title
// @Description
// @Author
// @Update

package utilities

import "fmt"

type LinkedList struct {
	head           *Node
	tail           *Node
	insertionOrder string
}

type Node struct {
	val  interface{}
	next *Node
}

const (
	HEAD = "Head"
	TAIL = "Tail"
)

func NewLinkedList(insertionOrder string) *LinkedList {
	return &LinkedList{insertionOrder: insertionOrder}
}

func (l *LinkedList) AddElement(val interface{}) interface{} {
	if l.insertionOrder == HEAD {
		return l.insertToHead(val)
	}

	return l.insertToTail(val)

}

func (l *LinkedList) DeleteElement(node *Node) {
	if node == nil {
		return
	}

	if node.GetNext() == nil {
		node = nil
		l.SetHead(nil)
		l.SetTail(nil)
		return
	}

	temp := node.GetNext()
	node.SetVal(temp.GetVal())
	node.SetNext(temp.GetNext())
	node = nil
}

// Prints the linked list
func (l *LinkedList) Print() {
	temp := l.GetHead()
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}

func (l *LinkedList) GetHead() *Node {
	return l.head
}

func (l *LinkedList) SetHead(head *Node) {
	l.head = head
}

func (l *LinkedList) GetTail() *Node {
	return l.tail
}

func (l *LinkedList) SetTail(tail *Node) {
	l.tail = tail
}

func (n *Node) GetVal() interface{} {
	return n.val
}

func (n *Node) SetVal(val interface{}) {
	n.val = val
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) SetNext(next *Node) {
	n.next = next
}

func (l *LinkedList) insertToHead(val interface{}) *Node {
	node := &Node{
		val:  val,
		next: nil,
	}
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}

	return node

}

func (l *LinkedList) insertToTail(val interface{}) *Node {
	node := &Node{
		val:  val,
		next: nil,
	}

	if l.head == nil {
		l.head = node
		l.tail = l.head
		return node
	}

	l.tail.next = node
	l.tail = node
	return node
}
