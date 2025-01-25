package main

func maxArea(height []int) int {
	largest := 0

	x, y := 0, len(height)-1
	for x < y {
		dist := y - x
		area := min(height[x], height[y]) * dist
		if area > largest {
			largest = area
		}

		if height[x] >= height[y] {
			y--
		} else {
			x++
		}
	}

	return largest
}
