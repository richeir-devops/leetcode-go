package P100

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := isBalanced_maxDepth(root.Left)
	rightHeight := isBalanced_maxDepth(root.Right)
	balance := leftHeight - rightHeight

	if balance <= 1 && balance >= -1 {
		return true
	} else {
		return false
	}

	balancedLeft := isBalanced(root.Left)
	balancedright := isBalanced(root.Right)

	return balancedLeft && balancedright
}

func isBalanced_maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left != nil {
		return 1 + isBalanced_maxDepth(root.Left)
	}

	if root.Left == nil && root.Right != nil {
		return 1 + isBalanced_maxDepth(root.Right)
	}

	maxLeft := 1 + isBalanced_maxDepth(root.Left)
	maxRight := 1 + isBalanced_maxDepth(root.Right)

	if maxLeft > maxRight {
		return maxLeft
	} else {
		return maxRight
	}
}

func Test_110_1(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		want bool
	}{
		{NewBinaryTree([]int{0, 3, 9, 20, -1, -1, 15, 7}), true},
	}

	for i, tt := range testCases {
		got := isBalanced(tt.root)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
