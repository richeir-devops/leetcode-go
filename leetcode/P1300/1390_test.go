package p1300

import (
	"testing"
)

// 想法：循环+模拟
// 注意极端值：只有一行，或者只有一列

func luckyNumbers(matrix [][]int) []int {
	var result []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			val := matrix[i][j]
			// 是否是一行中的最小值
			ismin := false
			for k := 0; k < len(matrix[i]); k++ {
				if k == j {
					continue
				}
				curval := matrix[i][k]
				if val > curval {
					ismin = false
					break
				}
				ismin = true
			}
			if len(matrix[i]) == 1 {
				ismin = true
			}
			if !ismin {
				continue
			}

			// 是否是一列中的最大值
			ismax := false
			for k := 0; k < len(matrix); k++ {
				if k == i {
					continue
				}
				curval := matrix[k][j]
				if val < curval {
					ismax = false
					break
				}
				ismax = true
			}
			if len(matrix) == 1 {
				ismax = true
			}
			if ismax {
				result = append(result, val)
			}
		}
	}

	return result
}

func Test_1380_01(t *testing.T) {
	t.Log("\n 1380. 矩阵中的幸运数 \n")
	matrix := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	t.Log(len(matrix))
	t.Log(len(matrix[0]))
	t.Log(matrix)
	t.Log(matrix[0])
	t.Log(luckyNumbers(matrix))
}

func Test_1380_02(t *testing.T) {
	t.Log("\n 1380. 矩阵中的幸运数 \n")
	matrix := [][]int{{56216}, {63251}, {75772}, {1945}, {27014}}
	t.Log(len(matrix))
	t.Log(len(matrix[0]))
	t.Log(matrix)
	t.Log(matrix[0])
	t.Log(luckyNumbers(matrix))
}

func Test_1380_03(t *testing.T) {
	t.Log("\n 1380. 矩阵中的幸运数 \n")
	matrix := [][]int{{76618, 42558, 65788, 20503, 29400, 54116}}
	t.Log(len(matrix))
	t.Log(len(matrix[0]))
	t.Log(matrix)
	t.Log(matrix[0])
	t.Log(luckyNumbers(matrix))
}
