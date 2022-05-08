package p1900

import "testing"

func minimumDifference(nums []int, k int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	bubbleSort(nums)

	var mnums []int
	for i := 0; i <= n-k; i++ {
		a := nums[i+k-1]
		b := nums[i]
		mnums = append(mnums, a-b)
	}

	bubbleSort(mnums)

	return mnums[0]
}

func bubbleSort(nums []int) {
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

func Test_1984_01(t *testing.T) {
	t.Log("\n 1984. 学生分数的最小差值 \n")
	nums := []int{9, 4, 1, 7}
	k := 2
	t.Log(minimumDifference(nums, k))
}

func Test_1984_02(t *testing.T) {
	t.Log("\n 1984. 学生分数的最小差值 \n")
	nums := []int{41900, 69441, 94407, 37498, 20299, 10856, 36221, 2231, 54526, 79072, 84309, 76765, 92282, 13401, 44698, 17586, 98455, 47895, 98889, 65298, 32271, 23801, 83153, 12186, 7453, 79460, 67209, 54576, 87785, 47738, 40750, 31265, 77990, 93502, 50364, 75098, 11712, 80013, 24193, 35209, 56300, 85735, 3590, 24858, 6780, 50086, 87549, 7413, 90444, 12284, 44970, 39274, 81201, 43353, 75808, 14508, 17389, 10313, 90055, 43102, 18659, 20802, 70315, 48843, 12273, 78876, 36638, 17051, 20478}
	k := 5
	t.Log(minimumDifference(nums, k))
}
