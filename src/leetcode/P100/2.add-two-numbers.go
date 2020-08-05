package leetcode

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 1
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

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

// @lc code=end
