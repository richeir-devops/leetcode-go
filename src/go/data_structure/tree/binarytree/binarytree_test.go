package normal

import "testing"

func Test_2293_01(t *testing.T) {
	testCases := []struct {
		treedata []int
		want     int
	}{
		{[]int{0, 3, 9, 20, -1, -1, 15, 7}, 3},
	}

	for i, tt := range testCases {

		got := NewBinaryTree(tt.treedata)
		if got.Val != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
