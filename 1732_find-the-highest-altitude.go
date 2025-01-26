package main

func largestAltitude(gain []int) int {
	curr, highest := 0, 0
	for _, alt := range gain {
		curr += alt
		if curr > highest {
			highest = curr
		}
	}

	return highest
}
