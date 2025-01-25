package main

func isSubsequence(s string, t string) bool {
	slow, fast := 0, 0
	for slow < len(s) && fast < len(t) {
		if t[fast] == s[slow] {
			slow += 1
		}

		fast++
	}

	// if we've transversed through the entire length
	// of s and all characters have been found
	// we can assume s is a subsequence of t
	return slow == len(s)
}
