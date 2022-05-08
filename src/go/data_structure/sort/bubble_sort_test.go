package sort

import "testing"

func bubble_sort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
}

func Test_01_bubble_sort(t *testing.T) {
	nums := []int{3, 4, 1, 5, 2, 1}
	t.Log(nums)
	bubble_sort(nums)
	t.Log(nums)
}
