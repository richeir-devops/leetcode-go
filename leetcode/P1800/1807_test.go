package p1800

import "testing"

func evaluate(s string, knowledge [][]string) string {
	result := ""
	srune := []rune(s)

	for i := 0; i < len(srune); i++ {
		if srune[i] == '(' {
			key := []rune{}
			for j := i + 1; j < len(srune); j++ {
				if srune[j] != ')' {
					key = append(key, srune[j])
				} else {
					break
				}
			}

			keyStr := string(key)
			i += len(keyStr) + 1
			var keyResult = evaluate_findkey(keyStr, knowledge)
			result += keyResult
		} else {
			result += string(srune[i])
		}
	}

	return result
}

func evaluate_findkey(key string, knowledge [][]string) string {
	for i := 0; i < len(knowledge); i++ {
		if knowledge[i][0] == key {
			return knowledge[i][1]
		}
	}

	return "?"
}

func Test_1807_01(t *testing.T) {
	testCases := []struct {
		s         string
		knowledge [][]string
		want      string
	}{
		{"(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}}, "bobistwoyearsold"},
	}

	for i, tt := range testCases {
		got := evaluate(tt.s, tt.knowledge)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
