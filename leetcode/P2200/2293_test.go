package p2200

import "testing"

func minMaxGame(nums []int) int {
	for {
		nlen := len(nums)
		if nlen == 1 {
			return nums[0]
		}

		newNums := []int{}
		for i := 0; i < nlen/2; i++ {
			if i%2 == 0 {
				if nums[2*i] < nums[2*i+1] {
					newNums = append(newNums, nums[2*i])
				} else {
					newNums = append(newNums, nums[2*i+1])
				}
			}

			if i%2 == 1 {
				if nums[2*i] > nums[2*i+1] {
					newNums = append(newNums, nums[2*i])
				} else {
					newNums = append(newNums, nums[2*i+1])
				}
			}
		}

		nums = newNums
	}
}

func Test_2293_01(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 5, 2, 4, 8, 2, 2}, 1},
		{[]int{3}, 3},
	}

	for i, tt := range testCases {
		got := minMaxGame(tt.nums)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
