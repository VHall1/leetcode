package main

func mergeAlternately(word1 string, word2 string) string {
	l1, l2 := len(word1), len(word2)
	merged := ""

	for i := 0; i < max(l1, l2); i++ {
		if i < l1 {
			merged = merged + string(word1[i])
		}
		if i < l2 {
			merged = merged + string(word2[i])
		}
	}

	return merged
}
