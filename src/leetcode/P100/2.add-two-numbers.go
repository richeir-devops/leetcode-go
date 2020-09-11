package p100

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

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var result = new(ListNode)
// 	result.Val = 0
// 	var plus = false
// 	var p = result
// 	for {
// 		if l1 != nil || l2 != nil {
// 			if l1 != nil {
// 				p.Val += l1.Val
// 				l1 = l1.Next
// 			}
// 			if l2 != nil {
// 				p.Val += l2.Val
// 				l2 = l2.Next
// 			}

// 			if p.Val > 9 {
// 				plus = true
// 				p.Val -= 10
// 			}

// 			if l1 != nil || l2 != nil || plus {
// 				var newNode = new(ListNode)
// 				if plus {
// 					newNode.Val++
// 					plus = false
// 				}
// 				p.Next = newNode
// 				p = newNode
// 			} else {
// 				break
// 			}
// 		} else {
// 			break
// 		}
// 	}

// 	plus = false
// 	return result
// }
