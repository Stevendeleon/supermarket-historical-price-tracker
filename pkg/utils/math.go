package utils

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// DELETE later this is just for testing pipeline purposes
