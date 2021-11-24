package SwordOffer

import (
	"testing"
)

func isNumber(s string) bool {
	states := map[int]map[rune]int{
		0: {
			' ': 0, 's': 1, 'd': 2, '.': 4,
		},
		1: {
			'd': 2, '.': 4,
		},
		2: {
			'd': 2, '.': 3, 'e': 5, ' ': 8,
		},
		3: {
			'd': 3, 'e': 5, ' ': 8,
		},
		4: {
			'd': 3,
		},
		5: {
			's': 6, 'd': 7,
		},
		6: {
			'd': 7,
		},
		7: {
			'd': 7, ' ': 8,
		},
		8: {
			' ': 8,
		},
	}

	p := 0
	var t rune
	for _, c := range s {
		if c >= '0' && c <= '9' {
			t = 'd'
		} else if c == '+' || c == '-' {
			t = 's'
		} else if c == 'e' || c == 'E' {
			t = 'e'
		} else if c == '.' || c == ' ' {
			t = c
		} else {
			t = '?'
		}
		_, found := states[p][t]
		if !found {
			return false
		}
		p = states[p][t]
	}

	return p == 2 || p == 3 || p == 7 || p == 8
}

func Test_020_01(t *testing.T) {
	t.Log(isNumber("123.."))
}
