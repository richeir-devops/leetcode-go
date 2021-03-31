package RicheirLinkedList

import "fmt"

// Element : 链表的结点内容描述
type Element struct {
	// next : 下一个节点
	next *Element
	// value : 当前节点的值，这里是 interface{} 类型
	value interface{}
}

// LinkedList : 这个对象包含根节点和链表长度
type LinkedList struct {
	// 根节点
	root Element
	len  int
}

// NewSingleLinkedList : 创建新的单链表实例
func NewSingleLinkedList() *LinkedList {
	// 通过 new(LinkedList) 来初始化一个 LinkedList 实例
	return new(LinkedList).init()
}

// init: 初始化一个单链表
func (l *LinkedList) init() *LinkedList {
	// 初始化的时候，根节点的 next 指向 nil
	l.root.next = nil
	// 链表长度初始化的时候是 0
	l.len = 0
	return l
}

// PushToHead : 从链表头部插入新的节点
func (l *LinkedList) PushToHead(insertValue interface{}) *Element {
	newNode := &Element{value: insertValue}

	// 注意前后顺序
	// 新节点的 next 指向根节点的 next
	newNode.next = l.root.next
	// 根节点的 next 指向新节点
	l.root.next = newNode

	l.len++
	return newNode
}

// PushToTail : 从尾部插入节点
func (l *LinkedList) PushToTail(insertValue interface{}) *Element {
	newNode := &Element{value: insertValue}

	// 找到尾节点（在没有单独定义尾节点的情况下，这也是大多数的情况）
	// 也可以单独维护一个尾节点来省略这个查找的步骤
	tailNode := l.root.next
	for i := 0; i < l.len-1; i++ {
		tailNode = tailNode.next
	}

	tailNode.next = newNode
	l.len++
	return newNode
}

// ReverseList : 反转当前链表
func (l *LinkedList) ReverseList() {
	var p, q *Element
	p = l.root.next
	q = nil
	//fireNode := true
	for p != nil {
		next := p.next
		p.next = q
		q = p
		if next == nil {
			l.root.next = p
			break
		}
		p = next
	}
}

func (l *LinkedList) PrintList() {
	node := l.root.next
	print("root")
	for i := 0; i < l.len; i++ {
		fmt.Printf(" -> %v", node.value)
		node = node.next
	}
	fmt.Printf("\n")
}
