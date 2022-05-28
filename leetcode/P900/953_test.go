package p900

import (
	"strings"
	"testing"
)

func isAlienSorted(words []string, order string) bool {
	wordsCopy := make([]string, len(words))
	copy(wordsCopy, words)
	bubble_sort_953(words, order)

	for i := 0; i < len(words); i++ {
		if wordsCopy[i] != words[i] {
			return false
		}
	}

	return true
}

func bubble_sort_953(nums []string, dict string) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if dict_compare_953(nums[j], nums[j+1], dict) {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
}

func dict_compare_953(a, b, dict string) bool {
	isNeedExchange := false
	lenA := len(a)
	lenB := len(b)
	lenMax := lenA
	if lenA < lenB {
		lenMax = lenB
	}

	for i := 1; i <= lenMax; i++ {
		if i <= lenA && i <= lenB {
			indexA := strings.Index(dict, string(a[i-1]))
			indexB := strings.Index(dict, string(b[i-1]))

			if indexA > indexB {
				return true
			} else if indexA < indexB {
				return false
			} else if indexA == indexB && i == lenA && lenA < lenB {
				return false
			} else if indexA == indexB && i == lenB && lenB < lenA {
				return true
			}
		}
	}

	return isNeedExchange
}

func Test_953_Mixed(t *testing.T) {

	testCases := []struct {
		words     []string
		dictIndex string
		want      bool
	}{
		{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
		{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{[]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz", true},
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"l", "h"}, "xkbwnqozvterhpjifgualycmds", false},
		{[]string{"app", "apple"}, "abcdefghijklmnopqrstuvwxyz", true},
	}

	for i, tt := range testCases {
		got := isAlienSorted(tt.words, tt.dictIndex)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \ngot: %v \nwant: %v \n\n", i, tt, got, tt.want)
		}
	}

	t.Log("All test cases are passed.")
}
