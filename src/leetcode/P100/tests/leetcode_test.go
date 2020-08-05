package leetcode

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func Test_1(t *testing.T) {
	a := []int{2, 7, 11, 15}
	t.Log(a)
	var result = twoSum(a, 9)
	if result[0] == 0 && result[1] == 1 {
		t.Log("OK")
	} else {
		t.Log("No")
	}
}

func Test_2(t *testing.T) {
	num := 10
	if num%2 == 0 { //checks if number is even
		t.Log("the number is even")
	} else {
		t.Log("the number is odd")
	}
}
