package SwordOffer

import (
	"testing"
)

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 每10个做一个采样
	sampling := []*ListNode{}
	lenList := 0
	for i := 1; head != nil; {
		if i%10 == 1 {
			sampling = append(sampling, head)
		}
		i++
		head = head.Next
		lenList++
	}

	// sampling[0] 表示 1
	// sampling[1] 表示 11
	// sampling[2] 表示 11

	//计算正数第几个
	pos := lenList - k + 1
	result := sampling[(pos-1)/10]
	for i := ((pos-1)/10)*10 + 1; i < pos; i++ {
		result = result.Next
	}

	return result
}

func Test_P022_1(t *testing.T) {
	t.Log("Sword Offer P022")
	list1 := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}
	getKthFromEnd(list1, 2)
}
