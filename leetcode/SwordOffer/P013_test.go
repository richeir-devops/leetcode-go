package SwordOffer

import (
	"testing"
)

func movingCount(m int, n int, k int) int {
	matrix := make([][]int, m)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	dfs013(&matrix, m, n, k, 0, 0)

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				result++
			}
		}
	}

	return result
}

func dfs013(matrix *[][]int, m, n, k, i, j int) {
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}

	if (*matrix)[i][j] == 1 {
		return
	}

	if p13CalcIntSum(i)+p13CalcIntSum(j) <= k {
		(*matrix)[i][j] = 1
	}

	if (*matrix)[i][j] == 1 {
		dfs013(matrix, m, n, k, i+1, j)
		dfs013(matrix, m, n, k, i, j+1)
	}
}

func p13CalcIntSum(num int) int {
	result := 0
	for num != 0 {
		result += num % 10
		num /= 10
	}

	return result
}

func Test_P013_01(t *testing.T) {
	m := 3
	n := 2
	matrix := make([][]int, m)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	t.Log(matrix)

	// t.Log(movingCount(2, 3, 1))
	t.Log(movingCount(16, 8, 4))
}
