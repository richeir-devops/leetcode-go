package P000

import "testing"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func Test001_1(t *testing.T) {
	t.Log("001")
	t.Log(twoSum([]int{2, 7, 11, 15}, 9))
}
