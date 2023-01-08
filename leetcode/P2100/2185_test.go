package P2100

import (
	"testing"
)

func prefixCount(words []string, pref string) int {
	cnt := 0
	hasPrefix := false
	for _, word := range words {
		hasPrefix = false
		if len(pref) > len(word) {
			continue
		}

		for i, p := range pref {
			if rune(word[i]) != p {
				break
			}
			if rune(word[i]) == p && i == (len(pref)-1) {
				hasPrefix = true
			}
		}

		if hasPrefix {
			cnt++
		}
	}

	return cnt
}

func Test_2185_01(t *testing.T) {
	testCases := []struct {
		words []string
		pref  string
		want  int
	}{
		{[]string{"pay", "attention", "practice", "attend"}, "at", 2},
		{[]string{"leetcode", "win", "loops", "success"}, "code", 0},
		{[]string{"cullp", "ypfyqcljk", "jtuogvscob", "dsriyigc", "fr", "jtuogvscou", "doivwcgxpo", "jtuogvscob", "chuiw", "fqdhyhkvtz", "cszrtrff", "kssjy", "ps", "xvax", "vekbmwwnra"}, "jtuogvsco", 3},
	}

	for i, tt := range testCases {
		got := prefixCount(tt.words, tt.pref)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
