// @Title
// @Description
// @Author
// @Update

package datastructure

type LinkedList struct {
	head           *Node
	tail           *Node
	insertionOrder string
}

type Node struct {
	val  interface{}
	next *Node
}

func NewLinkedList(insertionOrder string) LinkedList {
	return LinkedList{insertionOrder: insertionOrder}
}

func (l LinkedList) AddElement(val interface{}) interface{} {
	list := &l
	if l.insertionOrder == "Head" {
		return list.insertToHead(val)
	}

	return list.insertToTail(val)

}

func (l LinkedList) DeleteElement(val interface{}) {
	if l.head == nil {
		return
	}

	list := &l
	if val.(*Node) == l.head {
		temp := list.head
		list.head = temp.next
	} else {
		temp := list.head.next
		var prev *Node
		for temp != nil && val.(*Node) != temp {
			prev = temp
			temp = temp.next
		}

		prev.next.next = temp.next
		temp = nil

	}

}

func (l LinkedList) GetHead() *Node {
	return l.head
}

func (n *Node) GetVal() interface{} {
	return n.val
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

	l.tail.next = l.tail
	l.tail = node
	return node
}
