package p700

import "testing"

func toLowerCase(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if rs[i] >= 65 && rs[i] <= 90 {
			// 'a' - 'A' = 32
			rs[i] += 32
		}

	}

	return string(rs)
}

func Test_709_01(t *testing.T) {
	t.Log("\n 709. 转换成小写字母 \n")

	s := "Hello"
	t.Log(toLowerCase(s))
}
