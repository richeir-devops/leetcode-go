package mysinglelist

import "fmt"

// Element : list node
type Element struct {
	next *Element
	//list  *List
	Value interface{}
}

// List : root node of the list
type List struct {
	root Element
	tail *Element
	len  int
}

// Init initializes or clears list l
func (l *List) Init() *List {
	l.root.next = &l.root
	l.tail = &l.root
	l.len = 0
	return l
}

// NewSingleList returns an initialized list
func NewSingleList() *List {
	return new(List).Init()
}

// PushToHead insert
func (l *List) PushToHead(insertValue interface{}) *Element {
	newNode := new(Element)
	newNode.Value = insertValue

	newNode.next = l.root.next
	l.root.next = newNode

	l.len++
	return newNode
}

// PushToTail insert
func (l *List) PushToTail(insertValue interface{}) *Element {
	newNode := new(Element)
	newNode.Value = insertValue

	// find the tail node
	// tailNode := l.root.next
	// for i := 0; i < l.len-1; i++ {
	// 	tailNode = tailNode.next
	// }

	l.tail.next = newNode
	l.tail = newNode

	l.len++
	return newNode
}

// PrintList to the console
func (l *List) PrintList() {
	currentNode := l.root.next
	for i := 0; i < l.len; i++ {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.next
	}
}
