package mysinglelist

import "testing"

// Test1 11
func Test1(t *testing.T) {
	l1 := NewSingleList()
	if l1 != nil {
		t.Log("SingleList Created!")
	}

	l1.PushToHead(11)
	l1.PushToHead(12)
	l1.PrintList()
	t.Log("l1 OK!")

	l2 := NewSingleList()
	l2.PushToTail(1)
	l2.PushToTail(2)
	l2.PushToTail(3)
	l2.PushToTail(4)
	l2.PushToTail(5)
	l2.PrintList()
	t.Log("l2 OK!")
	l2.ReverseList()
	l2.PrintList()
}
