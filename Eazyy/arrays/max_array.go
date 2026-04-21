package arrays

func Max(arr [5]int) int {
	max := -999
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}
