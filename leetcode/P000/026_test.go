package P000

import "testing"

func removeDuplicates(nums []int) int {
	result := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] != nums[i] {
			for j := result; j <= i; j++ {
				nums[j] = nums[i+1]
			}
			result++
		}
	}

	return result
}

func Test_026_1(t *testing.T) {
	t.Log("026")
	t.Log(removeDuplicates([]int{1, 1, 2}))
}
