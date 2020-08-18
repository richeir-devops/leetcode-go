package daily202008

// import "testing"

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	leftBalance := false
// 	rightBalance := false
// 	leftLevel := 0
// 	rightLevel := 0
// 	if root.Left != nil {
// 		leftBalance, leftLevel = oneSide(root.Left)
// 		leftLevel++
// 	} else {
// 		leftBalance = true
// 	}

// 	if root.Right != nil {
// 		rightBalance, rightLevel = oneSide(root.Right)
// 		rightLevel++
// 	} else {
// 		rightBalance = true
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return true
// 	}

// 	levelDiffer := leftLevel - rightLevel

// 	if leftBalance && rightBalance {
// 		if levelDiffer >= -1 && levelDiffer <= 1 {
// 			return true
// 		}
// 	}

// 	return false
// }

// func oneSide(root *TreeNode) (bool, int) {
// 	if root == nil {
// 		return true, 0
// 	}

// 	leftLevel := 0
// 	leftCurrentLevel := 0
// 	rightLevel := 0
// 	rightCurrentLevel := 0

// 	if root.Left != nil {
// 		leftLevel++
// 		haveChild(root.Left, &leftCurrentLevel, &leftLevel)
// 	}

// 	if root.Right != nil {
// 		rightLevel++
// 		haveChild(root.Right, &rightCurrentLevel, &rightLevel)
// 	}

// 	levelDiffer := leftLevel - rightLevel

// 	totalLevel := 0
// 	if leftLevel < rightLevel {
// 		totalLevel = rightLevel
// 	} else {
// 		totalLevel = leftLevel
// 	}

// 	if levelDiffer < -1 || levelDiffer > 1 {
// 		return false, totalLevel
// 	}
// 	return true, totalLevel
// }

// func haveChild(node *TreeNode, currentLevel *int, totalLevel *int) {
// 	if *totalLevel > *currentLevel {
// 		*currentLevel++
// 	}

// 	hvaeTraveled := false

// 	if node.Left != nil {
// 		if !hvaeTraveled {
// 			*totalLevel++
// 			hvaeTraveled = true
// 		}
// 		haveChild(node.Left, currentLevel, totalLevel)
// 	}

// 	if node.Right != nil {
// 		if !hvaeTraveled {
// 			*totalLevel++
// 			hvaeTraveled = true
// 		}
// 		haveChild(node.Right, currentLevel, totalLevel)
// 	}
// }

// func Test_110_1(t *testing.T) {
// 	t.Log("110")
// 	b := 3
// 	var a *int
// 	a = &b
// 	*a++
// 	t.Log(b)

// 	// root := new(TreeNode)
// 	// root.Left = new(TreeNode)
// 	// root.Right = new(TreeNode)
// 	// root.Right.Left = new(TreeNode)
// 	// root.Right.Right = new(TreeNode)

// 	root := new(TreeNode)
// 	root.Right = new(TreeNode)
// 	root.Right.Right = new(TreeNode)

// 	t.Log(isBalanced(root))
// }
