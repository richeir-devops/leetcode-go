package p1800

import "testing"

func evaluate(s string, knowledge [][]string) string {
	result := ""
	srune := []rune(s)
	kmap := evaluate_get_mapped(knowledge)

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

			if keyResult, exist := kmap[keyStr]; exist {
				result += keyResult
			} else {
				result += "?"
			}

		} else {
			result += string(srune[i])
		}
	}

	return result
}

func evaluate_get_mapped(knowledge [][]string) map[string]string {
	kmap := make(map[string]string)
	for i := 0; i < len(knowledge); i++ {
		kmap[knowledge[i][0]] = knowledge[i][1]
	}

	return kmap
}

func Test_1807_01(t *testing.T) {
	testCases := []struct {
		s         string
		knowledge [][]string
		want      string
	}{
		{"(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}}, "bobistwoyearsold"},
		{"(name)is(age)yearsold", [][]string{{"name", "bob"}, {"agee", "two"}}, "bobis?yearsold"},
		{"hi(name)", [][]string{{"a", "b"}}, "hi?"},
		{"(z)h", [][]string{}, "?h"},
	}

	for i, tt := range testCases {
		got := evaluate(tt.s, tt.knowledge)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
