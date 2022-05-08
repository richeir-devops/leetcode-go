package SwordOffer

import (
	"testing"
)

// 解法1：使用深度有限搜索，就像贪吃蛇一样往前探索，然后把周围的都看一下
// 需要注意的是临界条件判定
// 另外数组传递不使用引用，因为要考虑到从矩阵中的任何一个位置开始去探索
func exist(board [][]byte, word string) bool {
	words := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, words, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word []byte, i int, j int, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 {
		return false
	}

	if board[i][j] != word[k] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	board[i][j] = '0'

	res := dfs(board, word, i+1, j, k+1) ||
		dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i, j+1, k+1) ||
		dfs(board, word, i, j-1, k+1)

	board[i][j] = word[k]

	return res
}

func Test_012_01(t *testing.T) {
	boards := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	t.Log(exist(boards, "ABCCED"))
}
