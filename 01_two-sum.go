package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		j, found := m[complement]

		if found {
			return []int{i, j}
		}

		m[num] = i
	}

	return []int{}
}
