package p1500

import (
	"strconv"
	"strings"
	"testing"
)

func reformatDate(date string) string {
	result := ""

	month := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	spiStr := strings.Split(date, " ")
	result += spiStr[2] + "-"
	for i := 1; i <= len(month); i++ {
		if month[i-1] == spiStr[1] {
			if i < 10 {
				result += "0" + strconv.Itoa(i) + "-"
			} else {
				result += strconv.Itoa(i) + "-"
			}
			break
		}
	}

	day := spiStr[0]
	dayNum := day[:len(day)-2]

	if len(dayNum) == 1 {
		result += "0"
	}
	result += dayNum

	return result
}

func Test_1500_01(t *testing.T) {
	testCases := []struct {
		date string
		want string
	}{
		{"20th Oct 2052", "2052-10-20"},
		{"6th Jun 1933", "1933-06-06"},
		{"26th May 1960", "1960-05-26"},
	}

	for i, tt := range testCases {
		got := reformatDate(tt.date)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \ngot: %v \nwant: %v \n\n", i, tt, got, tt.want)
		}
	}

	t.Log("All passed!")
}
