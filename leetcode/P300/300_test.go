package p300

import "testing"

// Not passed
func lexicalOrder(n int) []int {
	result := make([]int, 0)

	if n >= 10 {
		for i := 1; i <= n/10; i++ {
			result = append(result, i)
			for i := 0; i <= n%10; i++ {
				result = append(result, 10+i)
			}
		}

		n = (n % 10) - (n / 10)

		if n <= 0 {
			n = 2
		}

		for i := n; i < 10; i++ {
			result = append(result, i)
		}
	} else if n < 10 {
		for i := 1; i <= n; i++ {
			result = append(result, i)
		}
	}

	return result
}

func Test_300_01(t *testing.T) {
	t.Log("\n 386. 字典序排数 \n")
	t.Logf("%v \n", lexicalOrder(13))
}

func Test_300_02(t *testing.T) {
	t.Log("\n 386. 字典序排数 \n")
	t.Logf("%v \n", lexicalOrder(2))
}

func Test_300_03(t *testing.T) {
	t.Log("\n 386. 字典序排数 \n")
	t.Logf("%v \n", lexicalOrder(11))
}

func Test_300_04(t *testing.T) {
	t.Log("\n 386. 字典序排数 \n")
	t.Logf("%v \n", lexicalOrder(10))
}

func Test_300_05(t *testing.T) {
	t.Log("\n 386. 字典序排数 \n")
	t.Logf("%v \n", lexicalOrder(15))
}
