package P100

import "testing"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	return minDepth_caclDepth(root, 1)
}

func minDepth_caclDepth(root *TreeNode, deep int) int {
	deepLeft := deep
	deepRight := deep
	if root.Left != nil {
		deepLeft++
		deepLeft = minDepth_caclDepth(root.Left, deepLeft)
	}

	if root.Right != nil {
		deepRight++
		deepRight = minDepth_caclDepth(root.Right, deepRight)
	}

	if deepLeft <= deepRight {
		return deepLeft
	} else {
		return deepRight
	}
}

// === Test ===

func NewBinaryTree(treedata []int) *TreeNode {
	result := CreateTree(treedata, 1)
	return result
}

func CreateTree(treedata []int, pos int) *TreeNode {
	lent := len(treedata)
	if pos > lent || treedata[pos] == -1 {
		return nil
	}

	curNode := &TreeNode{Val: treedata[pos]}

	left := 2 * pos
	right := 2*pos + 1

	if left < len(treedata) {
		curNode.Left = CreateTree(treedata, left)
	}

	if right < len(treedata) {
		curNode.Right = CreateTree(treedata, right)
	}

	return curNode
}

func Test_2293_01(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		// {NewBinaryTree([]int{0, 3, 9, 20, -1, -1, 15, 7}), 2},
		{NewBinaryTree([]int{0, 2, -1, 3, -1, 4, -1, 5, -1, 6}), 5},
	}

	for i, tt := range testCases {
		got := minDepth(tt.root)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
