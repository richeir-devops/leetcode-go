package P000

import (
	"math"
	"strconv"
	"testing"
	"unsafe"
)

func reverse(x int) int {
	result := 0
	isPlus := true
	reverseInt := []int{}
	if x < 0 {
		isPlus = false
		x *= -1
	}

	for {
		reverseInt = append(reverseInt, x%10)
		x /= 10

		if x == 0 {
			break
		}
	}

	firstZero := true
	position10 := 1
	for i := len(reverseInt) - 1; i >= 0; i-- {
		if reverseInt[i] == 0 && firstZero {
			continue
		} else {
			firstZero = false
		}

		result += reverseInt[i] * position10
		position10 *= 10
	}

	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}

	if !isPlus {
		result *= -1
	}
	return result
}

func Test007_1(t *testing.T) {
	t.Log("007")

	a := int64(123)
	b := strconv.FormatInt(a, 10)
	t.Log(b)
	t.Log([]rune(b))

	t.Log(reverse(123))
	t.Logf("%v", unsafe.Sizeof(int(9646324351)))
}
