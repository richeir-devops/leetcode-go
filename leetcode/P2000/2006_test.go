package p2000

import "testing"

// 最简单粗暴，硬直接莽
func countKDifference(nums []int, k int) int {
	result := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			} else {
				n := nums[i] - nums[j]
				if n < 0 {
					n *= -1
				}
				if n == k {
					result++
				}
			}
		}
	}

	return result / 2
}

func Test_2006_1(t *testing.T) {
	t.Log("\n Leetcode 2006. 差的绝对值为 K 的数对数目 \n")
	arr := []int{1, 2, 2, 1}
	k := 1
	t.Log(countKDifference(arr, k))
}
