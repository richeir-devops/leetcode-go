package P000

import "testing"

func removeElement(nums []int, val int) int {
	result := len(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			removeElement_removeArrayElement(nums, i)
			result--
			i--
		}
	}

	return result
}

func removeElement_removeArrayElement(nums []int, index int) {
	for i := index; i < len(nums); i++ {
		if i == len(nums)-1 {
			nums[i] = -1
		} else {
			nums[i] = nums[i+1]
		}
	}
}

func Test_027_01(t *testing.T) {
	testCases := []struct {
		nums []int
		val  int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for i, tt := range testCases {
		got := removeElement(tt.nums, tt.val)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
