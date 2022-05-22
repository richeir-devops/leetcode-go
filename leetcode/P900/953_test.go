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

func Test_953_01(t *testing.T) {
	t.Logf("\n 953. 验证外星语词典 \n")

	words := []string{"apple", "app"}
	dictIndex := "abcdefghijklmnopqrstuvwxyz"

	// bubble_sort_953(words, dictIndex)
	// t.Log(words)

	t.Log(isAlienSorted(words, dictIndex)) // false

	words = []string{"word", "world", "row"}
	dictIndex = "worldabcefghijkmnpqstuvxyz"

	// bubble_sort_953(words, dictIndex)
	// t.Log(words)

	t.Log(isAlienSorted(words, dictIndex)) // false

}

func Test_953_02(t *testing.T) {
	words := []string{"kuvp", "q"}
	dictIndex := "ngxlkthsjuoqcpavbfdermiywz"

	t.Log(strings.Index(dictIndex, "k"))
	t.Log(strings.Index(dictIndex, "q"))

	t.Log(isAlienSorted(words, dictIndex)) // true
}

func Test_953_03(t *testing.T) {
	words := []string{"hello", "leetcode"}
	dictIndex := "hlabcdefgijkmnopqrstuvwxyz"

	t.Log(isAlienSorted(words, dictIndex)) // true
}

func Test_953_04(t *testing.T) {
	words := []string{"l", "h"}
	dictIndex := "xkbwnqozvterhpjifgualycmds"

	t.Log(isAlienSorted(words, dictIndex)) // false
}

func Test_953_05(t *testing.T) {
	words := []string{"app", "apple"}
	dictIndex := "abcdefghijklmnopqrstuvwxyz"

	t.Log(isAlienSorted(words, dictIndex)) // true
}
