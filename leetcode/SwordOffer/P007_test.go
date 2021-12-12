package SwordOffer

import "testing"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	dic := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		dic[inorder[i]] = i
	}
	return recur007(0, 0, len(inorder)-1, preorder, dic)
}

func recur007(root, left, right int, preorder []int, dic map[int]int) *TreeNode {
	if left > right {
		return nil
	}

	node := &TreeNode{Val: preorder[root]}
	i := dic[preorder[root]]
	node.Left = recur007(root+1, left, i-1, preorder, dic)
	node.Right = recur007(root+i-left+1, i+1, right, preorder, dic)

	return node
}

func Test_P007_01(t *testing.T) {

}
