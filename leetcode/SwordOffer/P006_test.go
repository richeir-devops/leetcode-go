package SwordOffer

import "testing"

//Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

///////////////////////////////////////////
// 解法1（普通逻辑）：
// 1.把单链表的值读取到一个数组里面
// 2.获取数组的长度，然后再从头到尾组成新的结果数组

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
	t.Log(reversePrint(list1))
}

// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：3.5 MB, 在所有 Go 提交中击败了 49.05% 的用户

///////////////////////////////////////////

///////////////////////////////////////////
// 解法2（回溯）：
// 1.终止条件: head == nil
// 2.递推工作：访问下一节点 head.next
// 3.回溯阶段：在当前结果集最后加上 head.Val
// []int 是普通数组
// &[]int 是数组的地址引用
// *[]int 可以作为参数接收 &[]int
// *[]int --> **[]int = []int

func reversePrint_2(head *ListNode) []int {
	result := make([]int, 0)
	return recur(head, &result)
}

func recur(head *ListNode, result *[]int) []int {
	if head == nil {
		return *result
	} else {
		recur(head.Next, result)
		*result = append(*result, head.Val)
	}

	return *result
}

func Test_P006_02(t *testing.T) {
	t.Log("Sword Offer P006")
	list1 := &ListNode{Val: 1,
		Next: &ListNode{Val: 3,
			Next: &ListNode{Val: 2},
		},
	}
	t.Log(reversePrint_2(list1))
}

// 执行用时：4 ms, 在所有 Go 提交中击败了 43.14% 的用户
// 内存消耗：4.6 MB, 在所有 Go 提交中击败了 38.27% 的用户

///////////////////////////////////////////
