package SwordOffer

import "testing"

// 方法1：递归（超时）
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return numWays(n-1) + numWays(n-2)
	}
}

// 方法2：动态规划，正向求解
func numWays2(n int) int {
	a := 1
	b := 1
	sum := 0
	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}

	return a
}

func Test_010_2_01(t *testing.T) {
	t.Log(numWays(7))
}

func Test_010_2_02(t *testing.T) {
	t.Log(numWays2(7))
}
