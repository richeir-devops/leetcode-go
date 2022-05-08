package p800

import "testing"

func shortestToChar(s string, c byte) []int {
	var lens = len(s)
	result := make([]int, lens)
	for i := 0; i < lens; i++ {
		result[i] = 99999
	}

	for i := 0; i < lens; i++ {
		if s[i] == c {
			result[i] = 0
		}
	}

	for i := 0; i < lens; i++ {
		if result[i] == 0 {
			for j := i - 1; j >= 0; j-- {
				dis := i - j
				if result[j] != 0 && dis < result[j] {
					result[j] = i - j
				} else {
					break
				}
			}

			for k := i + 1; k < lens; k++ {
				dis := k - i
				if result[k] != 0 && dis < result[k] {
					result[k] = k - i
				} else {
					break
				}
			}
		}
	}

	return result
}

func Test_821_01(t *testing.T) {
	t.Log("\n 821. 字符的最短距离 \n")
	t.Log(shortestToChar("loveleetcode", 'e'))
}

func Test_821_02(t *testing.T) {
	t.Log("\n 821. 字符的最短距离 \n")
	t.Log(shortestToChar("loveleetcode", 'e'))
}
