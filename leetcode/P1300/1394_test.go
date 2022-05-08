package p1300

import "testing"

func findLucky(arr []int) int {
	result := -1
	lenArr := make([]int, 501)

	for i := 0; i < len(arr); i++ {
		lenArr[arr[i]]++
	}

	for i := 1; i < len(lenArr); i++ {
		if lenArr[i] == i {
			result = i
		}
	}

	return result
}

func Test_1394_01(t *testing.T) {
	t.Log("\n 1394. 找出数组中的幸运数 \n")

	arr := []int{2, 2, 3, 4}

	t.Log(findLucky(arr))
}
