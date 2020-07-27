package singlelist

// Element : list node
type Element struct {
	next  *Element
	list  *List
	Value interface{}
}

// List : root node of the list
type List struct {
	root Element
	len  int
}

// Init initializes or clears list l
func (l *List) Init() *List {
	l.root.next = &l.root
	// l.root.prev = &l.root
	l.len = 0
	return l
}

// New returns an initialized list
func New() *List {
	return new(List).Init()
}
