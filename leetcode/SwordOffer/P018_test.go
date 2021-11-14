package SwordOffer

import "testing"

//Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		head = head.Next
		return head
	}

	var current *ListNode
	var pre *ListNode
	current = head
	pre = head
	for current != nil {
		if current.Val == val {
			pre.Next = current.Next
			break
		} else {
			pre = current
			current = current.Next
		}
	}

	return head
}

func Test_P018_01(t *testing.T) {
	t.Log("Sword Offer P006")
	list1 := &ListNode{Val: 4,
		Next: &ListNode{Val: 5,
			Next: &ListNode{Val: 1,
				Next: &ListNode{Val: 9},
			},
		},
	}
	deleteNode(list1, 5)
}
