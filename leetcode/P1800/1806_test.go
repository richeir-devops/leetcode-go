package p1800

import "testing"

func reinitializePermutation(n int) int {
	result := 0
	perm := make([]int, n)

	for i := 0; i < n; i++ {
		perm[i] = i
	}

	isInit := false

	for !isInit {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else if i%2 == 1 {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}

		perm = arr
		result++

		for i := 0; i < n; i++ {
			if perm[i] != i {
				break
			}

			if perm[i] == i && i == (n-1) {
				isInit = true
			}
		}
	}

	return result
}

func Test_1806_01(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		{2, 1},
		{4, 2},
		{6, 4},
	}

	for i, tt := range testCases {
		got := reinitializePermutation(tt.n)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
