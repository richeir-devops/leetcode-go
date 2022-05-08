package p1400

import "testing"

func destCity(paths [][]string) string {
	result := ""

	dict := make(map[string]string, 0)
	for i := 0; i < len(paths); i++ {
		dict[paths[i][0]] = paths[i][1]
	}

	result = paths[0][1]
	for i := 0; i < len(paths); i++ {
		if result == paths[i][0] {
			result = paths[i][1]
		} else {
			for j := 0; j < len(paths); j++ {
				if result == paths[j][0] {
					result = paths[j][1]
				}
			}
		}
	}

	return result
}

func Test_1463_01(t *testing.T) {
	t.Log("\n  1436. 旅行终点站 \n")

	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	t.Log(destCity(paths))

	paths = [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	t.Log(destCity(paths))
}
