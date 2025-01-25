package main

func lengthOfLongestSubstring(s string) int {
	cache := make([]int, 128)

	var start, largest int
	for end, char := range s {
		if cache[char] > start {
			start = cache[char]
		}

		// if curr window is larger than largest recorded
		if end-start+1 > largest {
			largest = end - start + 1
		}

		cache[char] = end + 1
	}

	return largest
}
