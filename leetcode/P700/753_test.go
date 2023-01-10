package p700

import "testing"

func crackSafe(n int, k int) string {
	return ""
}

func Test_753_01(t *testing.T) {
	testCases := []struct {
		n    int
		k    int
		want string
	}{
		{1, 2, "01"},
		{2, 2, "00110"},
	}

	for i, tt := range testCases {
		got := crackSafe(tt.n, tt.k)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
