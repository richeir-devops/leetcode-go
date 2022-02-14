package p500

import (
	"testing"
)

// 想法：用穷举的方式来解答：
// 1. len = 1
// 2. len >= 3
// 3. 第一个元素是答案
// 4. 最后一个元素是答案
// 5. 循环起始是第二个到倒数第二个

func singleNonDuplicate(nums []int) int {
	result := 0

	// 情况1
	if len(nums) == 1 {
		return nums[0]
	}

	// 情况2和3
	if len(nums) >= 3 {
		if nums[0] != nums[1] && nums[1] == nums[2] {
			return nums[0]
		}
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			result = nums[i]
		}
	}

	// 情况4
	if result == 0 {
		result = nums[len(nums)-1]
	}

	return result
}

func Test_540_01(t *testing.T) {
	t.Log("\n 540. 有序数组中的单一元素 \n")
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	t.Log(singleNonDuplicate(nums))
}

func Test_540_02(t *testing.T) {
	t.Log("\n 540. 有序数组中的单一元素 \n")
	nums := []int{1, 1, 2}
	t.Log(singleNonDuplicate(nums))
}

func Test_540_03(t *testing.T) {
	t.Log("\n 540. 有序数组中的单一元素 \n")
	nums := []int{1, 2, 2, 3, 3}
	t.Log(singleNonDuplicate(nums))
}
