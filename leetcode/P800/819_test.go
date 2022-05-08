package p800

import (
	"strings"
	"testing"
)

func mostCommonWord(paragraph string, banned []string) string {
	result := ""
	// !?',;.
	paragraph = strings.ReplaceAll(paragraph, "!", " ")
	paragraph = strings.ReplaceAll(paragraph, "?", " ")
	paragraph = strings.ReplaceAll(paragraph, "'", " ")
	paragraph = strings.ReplaceAll(paragraph, ",", " ")
	paragraph = strings.ReplaceAll(paragraph, ";", " ")
	paragraph = strings.ReplaceAll(paragraph, ".", " ")
	paragraph = strings.ReplaceAll(paragraph, "  ", " ")

	splitedArrs := strings.Split(paragraph, " ")

	countMap := make(map[string]int)

	for i := 0; i < len(splitedArrs); i++ {
		var lowerItem = strings.ToLower(splitedArrs[i])
		countMap[lowerItem]++
	}

	max := 0
	for key, value := range countMap {
		isBanned := false

		if key == "" || key == " " {
			continue
		}

		for j := 0; j < len(banned); j++ {
			if key == banned[j] {
				isBanned = true
			}
		}

		if !isBanned && value > max {
			max = value
			result = key
		}
	}

	return result
}

func Test_819_01(t *testing.T) {
	t.Log("\n 819. 最常见的单词 \n")
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}

	t.Logf("%q\n", strings.Split(paragraph, " "))
	t.Logf("%q\n", strings.ReplaceAll(paragraph, ",", ""))
	t.Logf("%q\n", strings.ReplaceAll(paragraph, ".", ""))

	t.Log(mostCommonWord(paragraph, banned))
}

func Test_819_02(t *testing.T) {
	t.Log("\n 819. 最常见的单词 \n")
	paragraph := "a."
	banned := []string{}

	t.Log(mostCommonWord(paragraph, banned))
}

func Test_819_03(t *testing.T) {
	t.Log("\n 819. 最常见的单词 \n")
	paragraph := "a, a, a, a, b,b,b,c, c"
	banned := []string{"a"}

	t.Log(mostCommonWord(paragraph, banned))
}

func Test_819_04(t *testing.T) {
	t.Log("\n 819. 最常见的单词 \n")
	paragraph := "Bob!"
	banned := []string{"hit"}

	t.Log(mostCommonWord(paragraph, banned))
}

func Test_819_05(t *testing.T) {
	t.Log("\n 819. 最常见的单词 \n")
	paragraph := "abc abc? abcd the jeff!"
	banned := []string{"abc", "abcd", "jeff"}

	t.Log(mostCommonWord(paragraph, banned))
}
