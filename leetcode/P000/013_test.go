package P000

import "testing"

func romanToInt(s string) int {
	result := 0
	lenS := len(s) - 1
	for i := 0; i <= lenS; i++ {
		var currentLetter = s[i]
		switch currentLetter {

		case 'I':
			if lenS-i >= 1 {
				if s[i+1] == 'V' {
					result += 4
					i++
				} else if s[i+1] == 'X' {
					result += 9
					i++
				} else {
					result++
				}
			} else {
				result++
			}

		case 'V':
			result += 5

		case 'X':
			if lenS-i >= 1 {
				if s[i+1] == 'L' {
					result += 40
					i++
				} else if s[i+1] == 'C' {
					result += 90
					i++
				} else {
					result += 10
				}
			} else {
				result += 10
			}

		case 'L':
			result += 50

		case 'C':
			if lenS-i >= 1 {
				if s[i+1] == 'D' {
					result += 400
					i++
				} else if s[i+1] == 'M' {
					result += 900
					i++
				} else {
					result += 100
				}
			} else {
				result += 100
			}

		case 'D':
			result += 500

		case 'M':
			result += 1000
		}
	}

	return result
}

func Test_013_1(t *testing.T) {
	t.Log("013")
	a := "III"
	t.Log(a[1])
	t.Log(a[1] == 'I')

	t.Log(romanToInt("III"))
	t.Log(romanToInt("MCMXCIV"))
	t.Log(romanToInt("DCXXI"))
}
