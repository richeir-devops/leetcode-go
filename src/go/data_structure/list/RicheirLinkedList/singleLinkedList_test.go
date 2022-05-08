package RicheirLinkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Init_01(t *testing.T) {
	list1 := NewSingleLinkedList()
	// 写断言
	a := assert.New(t)
	a.Equal(0, list1.len)

	if list1.root.next == nil {
		t.Log("list1.root.next is nil")
	}
	// 这里不相等，很奇怪，明明 if 判断是过了的
	a.Equal(nil, list1.root.next)
}

func Test_Insert_Node_01(t *testing.T) {
	list1 := NewSingleLinkedList()
	list1.PushToHead(1)
	list1.PushToHead(2)
	list1.PushToHead(3)
	list1.PushToHead(4)
	list1.PushToHead(5)
	list1.PrintList()
	list1.ReverseList()
	list1.PrintList()
}

func Test_Insert_Node_02(t *testing.T) {
	list1 := NewSingleLinkedList()
	list1.PushToTail(1)
	list1.PushToTail(2)
	list1.PushToTail(3)
	list1.PushToTail(4)
	list1.PushToTail(5)
	list1.PrintList()
	list1.ReverseList()
	list1.PrintList()
}
