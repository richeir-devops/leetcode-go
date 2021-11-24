package SwordOffer

import "testing"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法1：迭代（双指针）
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

// 方法2：递归
func reverseList2(head *ListNode) *ListNode {
	return recur2(head, nil)
}

func recur2(cur *ListNode, pre *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	res := recur2(cur.Next, cur)
	cur.Next = pre
	return res
}

func Test_P024_01(t *testing.T) {
	t.Log("Sword Offer P024")
	list1 := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3},
		},
	}
	t.Log(reverseList(list1))
}

func Test_P024_02(t *testing.T) {
	t.Log("Sword Offer P024")
	list1 := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3},
		},
	}
	t.Log(reverseList2(list1))
}
