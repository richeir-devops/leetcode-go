package P000

import (
	"strings"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	prefix := ""

	if len(strs) == 0 {
		return prefix
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for i := 1; i <= len(strs[0]); i++ {
		prefix = strs[0][0:i]
		allPass := true
		for j := 0; j < len(strs); j++ {
			if allPass && strings.HasPrefix(strs[j], prefix) {

			} else {
				allPass = false
				if len(prefix) > 1 {
					prefix = strs[0][0 : i-1]
				} else {
					prefix = ""
				}

				return prefix
			}
		}
	}

	return prefix
}

func Test_014_1(t *testing.T) {
	t.Log("014")
	a := "123456"
	t.Log(a[1:3])
	t.Log(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	t.Log(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	t.Log(longestCommonPrefix([]string{}))
}
