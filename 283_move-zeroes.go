package main

func moveZeroes(nums []int) {
	lastNonZero := 0

	// move non-zero integers to the front and
	// save pointer to last non-zero integer
	// so we know where to fill with zeroes later
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZero] = nums[i]
			lastNonZero += 1
		}
	}

	// fill the rest of the array with zeroes
	for i := lastNonZero; i < len(nums); i++ {
		nums[i] = 0
	}
}
