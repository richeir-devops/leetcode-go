package SwordOffer

import (
	"testing"
)

// 剑指 Offer 10- I. 斐波那契数列

// 方法1：标准递归（会超时）
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

// 方法2：记忆化递归（也会超时）
func fib2(n int) int {
	fibMap := make(map[int]int)

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		if v, found := findInMap(&fibMap, n); found {
			return v
		} else {
			return fib2(n-1) + fib2(n-2)
		}
	}
}

func findInMap(fibMap *map[int]int, n int) (int, bool) {
	if v, found := (*fibMap)[n]; found {
		return v, true
	} else {
		return 0, false
	}
}

// 方法3：动态规划（实际上市正向求解）
func fib3(n int) int {
	a := 0
	b := 1
	sum := 0
	for i := 0; i < n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}

	return a
}

func Test_010_1(t *testing.T) {
	t.Log(fib(10))
}

func Test_010_2(t *testing.T) {
	t.Log(fib(10))
}
