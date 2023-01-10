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

	balance, _ := CheckNode(root)

	if balance {
		return true
	}

	return false
}

func CheckNode(node *TreeNode) (bool, int) {
	deepL := 0
	deepR := 0
	allBalance := true

	if node.Left != nil {
		allBalance, deepL = CheckNode(node.Left)
		if !allBalance {
			return false, 0
		}
	}

	if node.Right != nil {
		allBalance, deepR = CheckNode(node.Right)
		if !allBalance {
			return false, 0
		}
	}

	//叶子节点处理
	if node.Left == nil && node.Right == nil {
		return true, 1
	}

	//判断当前节点的是否平衡
	if node.Left != nil && node.Right != nil {
		// deepL++
		// deepR++

		if deepL-deepR <= 1 && deepL-deepR >= -1 {
			allBalance = true
		} else {
			return false, 0
		}

		if deepL > deepR {
			return true, deepL + 1
		}
		return true, deepR + 1
	} else if node.Left != nil && node.Right == nil {
		// deepL++

		if deepL-deepR <= 1 && deepL-deepR >= -1 {
			allBalance = true
		} else {
			return false, 0
		}

		return allBalance, deepL + 1
	} else {
		// deepR++

		if deepL-deepR <= 1 && deepL-deepR >= -1 {
			allBalance = true
		} else {
			return false, 0
		}

		return allBalance, deepR + 1
	}
}

func Test_110_1(t *testing.T) {
	t.Log("110")
	b := 3
	var a *int
	a = &b
	*a++
	t.Log(b)

	// root := new(TreeNode)
	// root.Left = new(TreeNode)
	// root.Right = new(TreeNode)
	// root.Right.Left = new(TreeNode)
	// root.Right.Right = new(TreeNode)

	root := new(TreeNode)
	root.Right = new(TreeNode)
	root.Right.Right = new(TreeNode)

	// root := new(TreeNode)
	// root.Left = new(TreeNode)

	t.Log(isBalanced(root))

	ch := 'b'
	t.Log(ch / 2.0)
}
