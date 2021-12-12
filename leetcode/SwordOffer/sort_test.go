package SwordOffer

import "testing"

func quick_sort(nums []int, l, r int) {
	if l >= r {
		return
	}

	i := partition(nums, l, r)
	quick_sort(nums, l, i-1)
	quick_sort(nums, i+1, r)
}

func partition(nums []int, l, r int) int {
	i := l
	j := r
	for i < j {
		for i < j && nums[j] >= nums[l] {
			j--
		}

		for i < j && nums[i] <= nums[l] {
			i++
		}
		swap(nums, i, j)
	}
	swap(nums, i, l)
	return i
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func Test_quick_sort(t *testing.T) {
	nums := []int{4, 1, 3, 2, 5}
	quick_sort(nums, 0, len(nums)-1)
	t.Log(nums)
}
