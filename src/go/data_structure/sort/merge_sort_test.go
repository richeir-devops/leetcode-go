package sort

import "testing"

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	m := (l + r) / 2
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, r)

	tmp := make([]int, r-l+1)
	for k := l; k <= r; k++ {
		tmp[k-l] = nums[k]
	}

	i := 0
	j := m - l + 1
	for k := l; k <= r; k++ {
		if i == m-l+1 {
			nums[k] = tmp[j]
			j++
		} else if j == r-l+1 || tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}
}

func Test01_merge_sort(t *testing.T) {
	nums := []int{3, 4, 1, 5, 2, 1}
	t.Log(nums)
	mergeSort(nums, 0, len(nums)-1)
	t.Log(nums)
}
