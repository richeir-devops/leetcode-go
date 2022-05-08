package daily202008

import "testing"

func judgeCircle(moves string) bool {
	result := false

	s1 := []rune(moves)
	x, y := 0, 0
	for i := 0; i < len(s1); i++ {
		makeStep(&x, &y, s1[i])
	}

	if x == 0 && y == 0 {
		result = true
	}

	return result
}

func makeStep(nowX *int, nowY *int, toward rune) {
	switch toward {
	case 85:
		//U
		*nowY++
	case 68:
		//D
		*nowY--
	case 76:
		//L
		*nowX--
	case 82:
		*nowX++
	}
}

func Test_657_1(t *testing.T) {
	t.Log("657")
	s1 := "UD"
	// s2 := []rune(s1)
	// t.Log(s2[0])

	t.Log(judgeCircle(s1))
}
