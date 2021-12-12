package SwordOffer

import "testing"

func getLeastNumbers(arr []int, k int) []int {
	quick_sort(arr, 0, len(arr)-1)
	return arr[:k]
}

func Test_P040_1(t *testing.T) {
	arr := []int{0, 1, 2, 1}
	t.Log(getLeastNumbers(arr, 1))
}
