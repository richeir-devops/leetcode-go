package leetcode

import (
	"fmt"
	"testing"
)

// ListNode : list node
// type ListNode struct {
// 	Next *ListNode
// 	//list  *List
// 	Value interface{}
// }

// List : root node of the list
type List struct {
	root ListNode
	tail *ListNode
	len  int
}

// Init initializes or clears list l
func (l *List) Init() *List {
	l.root.Next = &l.root
	l.tail = &l.root
	l.len = 0
	return l
}

// NewSingleList returns an initialized list
func NewSingleList() *List {
	return new(List).Init()
}

// PushToHead insert
func (l *List) PushToHead(insertValue int) *ListNode {
	newNode := new(ListNode)
	newNode.Val = insertValue

	newNode.Next = l.root.Next
	l.root.Next = newNode

	l.len++
	return newNode
}

// PushToTail insert
func (l *List) PushToTail(insertValue int) *ListNode {
	newNode := new(ListNode)
	newNode.Val = insertValue

	// find the tail node
	// tailNode := l.root.Next
	// for i := 0; i < l.len-1; i++ {
	// 	tailNode = tailNode.Next
	// }

	l.tail.Next = newNode
	l.tail = newNode

	l.len++
	return newNode
}

// PrintList to the console
func (l *List) PrintList() {
	currentNode := l.root.Next
	for i := 0; i < l.len; i++ {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}
}

// ListNode 1
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedListToNumber(list *ListNode) int {
	var result = 0
	var i = 1
	for {
		if list != nil {
			result += list.Val * i
			i *= 10
			list = list.Next
		} else {
			break
		}
	}

	return result
}

func numToReverseLinkedList(num int) *ListNode {
	var head *ListNode
	var p *ListNode
	for {
		var node = new(ListNode)
		node.Val = num % 10
		if head == nil {
			head = node
			p = head
		} else {
			p.Next = node
			p = p.Next
		}

		num /= 10

		if num == 0 {
			break
		}
	}

	return head
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var num1 = reverseLinkedListToNumber(l1)
// 	var num2 = reverseLinkedListToNumber(l2)
// 	var result = num1 + num2

// 	return numToReverseLinkedList(result)
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = new(ListNode)
	result.Val = 0
	var plus = false
	var p = result
	for {
		if l1 != nil || l2 != nil {
			if l1 != nil {
				p.Val += l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				p.Val += l2.Val
				l2 = l2.Next
			}

			if p.Val > 9 {
				plus = true
				p.Val -= 10
			}

			if l1 != nil || l2 != nil || plus {
				var newNode = new(ListNode)
				if plus {
					newNode.Val++
					plus = false
				}
				p.Next = newNode
				p = newNode
			} else {
				break
			}
		} else {
			break
		}
	}

	plus = false
	return result
}

func Test_002_1(t *testing.T) {
	l1 := NewSingleList()
	l1.PushToTail(2)
	l1.PushToTail(4)
	l1.PushToTail(3)
	l1.PrintList()

	t.Log("====")

	l2 := NewSingleList()
	l2.PushToTail(5)
	l2.PushToTail(6)
	l2.PushToTail(4)
	l2.PrintList()

	var result = addTwoNumbers(l1.root.Next, l2.root.Next)
	result.Val = 1
}
