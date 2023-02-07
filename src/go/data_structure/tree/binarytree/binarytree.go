package normal

const MagicNodeNilValue int = -1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBinaryTree(treedata []int) *TreeNode {
	result := CreateTree(treedata, 1)
	return result
}

func CreateTree(treedata []int, pos int) *TreeNode {
	lent := len(treedata)
	if pos > lent || treedata[pos] == MagicNodeNilValue {
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
