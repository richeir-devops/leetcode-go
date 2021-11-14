package SwordOffer

import "testing"

//Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func reversePrint(head *ListNode) []int {
	result := []int{}
	tempSlice := []int{}
	for i := 0; head != nil; i++ {
		tempSlice = append(tempSlice, head.Val)
		head = head.Next
	}

	lenSlice := len(tempSlice)
	for i := lenSlice - 1; i >= 0; i-- {
		result = append(result, tempSlice[i])
	}

	return result
}

func Test_P006_01(t *testing.T) {
	t.Log("Sword Offer P006")
	list1 := &ListNode{Val: 1,
		Next: &ListNode{Val: 3,
			Next: &ListNode{Val: 2},
		},
	}
	reversePrint(list1)
}
