package daily202008

import "testing"

func letterCombinations(digits string) []string {
	var result []string

	if len(digits) == 0 {
		return result
	}

	temp := make([]byte, 0)
	digui17(&result, digits, 0, temp)

	return result
}

func getLetter(num int) string {
	num2 := "abc"
	num3 := "def"
	num4 := "ghi"
	num5 := "jkl"
	num6 := "mno"
	num7 := "pqrs"
	num8 := "tuv"
	num9 := "wxyz"

	switch num {
	case 2:
		return num2
	case 3:
		return num3
	case 4:
		return num4
	case 5:
		return num5
	case 6:
		return num6
	case 7:
		return num7
	case 8:
		return num8
	case 9:
		return num9
	}

	return ""
}

func digui17(result *[]string, targetStr string, currenStrPos int, currentStr []byte) {
	//当前处理的位置如果正好等于最大长度，则说明已经得到一条有效的数据，否则继续
	if len(currentStr) == len(targetStr) {
		*result = append(*result, string(currentStr))
	} else {
		var letters = getLetter(int(targetStr[currenStrPos]) - 48)
		currenStrPos++
		for i := 0; i < len(letters); i++ {
			tempStrBytes := append(currentStr, letters[i])
			digui17(result, targetStr, currenStrPos, tempStrBytes)
		}
	}
}

func Test_17_1(t *testing.T) {
	t.Log("17")

	s1 := "23"
	t.Log(s1[1] - 48)

	var result = make([]string, 0)
	// var d1 = &result
	// var d2 = *d1
	// *d1 = append(*d1, "321")

	// result = append(result, "aa")
	// result = append(result, "ab")

	temp := make([]byte, 0)
	digui17(&result, "23", 0, temp)

	t.Log(result)
}
