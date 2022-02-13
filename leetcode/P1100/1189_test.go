package p1900

import (
	"math"
	"testing"
)

// 想法：遍历字符串，找到对应的 b,a,l,o,n 的字符的熟练
// 然后按照 balloon 单词至少需要 1个b，1个a，2个l，2个o，1个n 来判断能组出来几个单词，就像木桶原理取 最小的值/2 的即可
// 0/2 和 1/2 都是不做四舍五入的，结果都是 0

func maxNumberOfBalloons(text string) int {
	n := len(text)
	letterArray := make([]int, 5)

	for i := 0; i < n; i++ {
		switch text[i] {
		case 'b':
			letterArray[0] += 2
		case 'a':
			letterArray[1] += 2
		case 'l':
			letterArray[2] += 1
		case 'o':
			letterArray[3] += 1
		case 'n':
			letterArray[4] += 2
		}
	}

	min := math.MaxInt
	for i := 0; i < len(letterArray); i++ {
		if min > letterArray[i] {
			min = letterArray[i]
		}
	}

	return min / 2
}

func Test_1189_01(t *testing.T) {
	t.Log("\n 1189. “气球” 的最大数量 \n")
	text := "nlaebolko"
	t.Log(maxNumberOfBalloons(text))
}
