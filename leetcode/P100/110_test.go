package P100

import (
	"testing"
)

func isBalanced(root *TreeNode) bool {
	return isBalanced_Height(root) >= 0
}

func isBalanced_Height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := isBalanced_Height(root.Left)
	rightHeight := isBalanced_Height(root.Right)

	balance := leftHeight - rightHeight
	isBalance := false
	if balance <= 1 && balance >= -1 {
		isBalance = true
	}

	if leftHeight >= 0 && rightHeight >= 0 && isBalance {
		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
	} else {
		return -1
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
