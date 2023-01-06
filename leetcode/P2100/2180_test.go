package P2100

import "testing"

func countEven(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if countEven_isEven(i) {
			count++
		}
	}

	return count
}

func countEven_isEven(num int) bool {
	sum := 0
	for true {
		sum += num % 10
		num /= 10
		if num == 0 {
			break
		}
	}

	return sum%2 == 0
}

func Test_2180_01(t *testing.T) {
	testCases := []struct {
		num  int
		want int
	}{
		{4, 2},
		{30, 14},
	}

	for i, tt := range testCases {
		got := countEven(tt.num)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
