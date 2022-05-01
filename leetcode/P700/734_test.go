package p700

import "testing"

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {

	// #1 判断句子长度
	if len(sentence1) != len(sentence2) {
		return false
	}

	// #2 逐个单词判断
	for i := 0; i < len(sentence1); i++ {
		s1 := sentence1[i]
		s2 := sentence2[i]
		if s1 == s2 {
			continue
		}

		if s1 != s2 {
			isPair := false
			for j := 0; j < len(similarPairs); j++ {
				sp0 := similarPairs[j][0]
				sp1 := similarPairs[j][1]
				if (sp0 == s1 || sp0 == s2) && (sp1 == s1 || sp1 == s2) {
					isPair = true
					break
				}
			}
			if !isPair {
				return false
			}
		}
	}

	return true
}

func Test_734_01(t *testing.T) {
	t.Log("\n 734. 句子相似性 \n")

	sentence1 := []string{"great"}
	sentence2 := []string{"great"}
	similarPairs := [][]string{}
	t.Log(areSentencesSimilar(sentence1, sentence2, similarPairs))

	sentence1 = []string{"great", "acting", "skills"}
	sentence2 = []string{"fine", "drama", "talent"}
	similarPairs = [][]string{{"great", "fine"}, {"drama", "acting"}, {"skills", "talent"}}

	t.Log(areSentencesSimilar(sentence1, sentence2, similarPairs))

	sentence1 = []string{"great"}
	sentence2 = []string{"doubleplus", "good"}
	similarPairs = [][]string{{"great", "doubleplus"}}

	t.Log(areSentencesSimilar(sentence1, sentence2, similarPairs))
}

func Test_734_02(t *testing.T) {
	sentence1 := []string{"an", "extraordinary", "meal"}
	sentence2 := []string{"one", "good", "dinner"}
	similarPairs := [][]string{{"great", "good"}, {"extraordinary", "good"}, {"well", "good"}, {"wonderful", "good"}, {"excellent", "good"}, {"fine", "good"}, {"nice", "good"}, {"any", "one"}, {"some", "one"}, {"unique", "one"}, {"the", "one"}, {"an", "one"}, {"single", "one"}, {"a", "one"}, {"truck", "car"}, {"wagon", "car"}, {"automobile", "car"}, {"auto", "car"}, {"vehicle", "car"}, {"entertain", "have"}, {"drink", "have"}, {"eat", "have"}, {"take", "have"}, {"fruits", "meal"}, {"brunch", "meal"}, {"breakfast", "meal"}, {"food", "meal"}, {"dinner", "meal"}, {"super", "meal"}, {"lunch", "meal"}, {"possess", "own"}, {"keep", "own"}, {"have", "own"}, {"extremely", "very"}, {"actually", "very"}, {"really", "very"}, {"super", "very"}}
	// should be true
	t.Log(areSentencesSimilar(sentence1, sentence2, similarPairs))
}
