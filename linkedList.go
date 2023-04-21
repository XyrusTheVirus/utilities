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

func (l *LinkedList) DeleteElement(val interface{}) {
	if l.head == nil {
		return
	}

	if val.(*Node) == l.head {
		temp := l.head
		l.head = temp.next
	} else {
		temp := l.head.next
		var prev *Node
		for temp != nil && val.(*Node) != temp {
			prev = temp
			temp = temp.next
		}

		prev.next.next = temp.next
		temp = nil

	}

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
