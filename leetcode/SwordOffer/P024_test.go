package SwordOffer

import "testing"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var current, pre *ListNode
	current = head
	pre = nil
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		if next == nil {
			head = current
			break
		}
		current = next
	}
	return head
}

func Test_P024_01(t *testing.T) {
	t.Log("Sword Offer P024")
	list1 := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3},
		},
	}
	reverseList(list1)
}
