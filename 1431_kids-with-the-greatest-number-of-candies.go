package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	out := make([]bool, len(candies))

	max := 0
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		out[i] = candies[i]+extraCandies >= max
	}

	return out
}
