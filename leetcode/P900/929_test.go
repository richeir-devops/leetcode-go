package p900

import (
	"strings"
	"testing"
)

func numUniqueEmails(emails []string) int {
	targetMap := make(map[string]int)

	for _, v := range emails {
		name := v[:strings.IndexAny(v, "@")]
		domain := v[strings.IndexAny(v, "@"):]

		if strings.IndexAny(name, ".") > 0 {
			name = strings.ReplaceAll(name, ".", "")
		}
		if strings.IndexAny(name, "+") > 0 {
			name = name[:strings.Index(name, "+")]
		}

		v = name + domain
		targetMap[v]++
	}

	return len(targetMap)
}

func Test_929(t *testing.T) {
	testCases := []struct {
		emails []string
		want   int
	}{
		{[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2},
		{[]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}, 3},
		{[]string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}, 1},
		{[]string{"linqmafia@leet+code.com", "linqmafia@code.com"}, 2},
	}

	for i, tt := range testCases {
		got := numUniqueEmails(tt.emails)
		if got != tt.want {
			t.Errorf("\nTest Case #%v \ndata: %v \nwant: %v \ngot: %v \n\n", i, tt, tt.want, got)
		}
	}
}
