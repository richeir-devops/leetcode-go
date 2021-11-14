package SwordOffer

import "testing"

func replaceSpace(s string) string {
	intStr := []rune(s)
	lenStr := len(intStr)
	count := 0

	for i := 0; i < lenStr; i++ {
		if intStr[i] == rune(' ') {
			count++
		}
	}

	intStr = append(intStr, make([]rune, count*2)...)
	j := len(intStr) - 1
	for i := lenStr - 1; i < j; i-- {
		if intStr[i] != rune(' ') {
			intStr[j] = intStr[i]
		} else {
			intStr[j-2] = rune('%')
			intStr[j-1] = rune('2')
			intStr[j] = rune('0')
			j -= 2
		}
		j--
	}

	return string(intStr)
}

func Test_P005_1(t *testing.T) {
	input := "We are happy."
	output := replaceSpace(input)
	t.Log(output)
}
