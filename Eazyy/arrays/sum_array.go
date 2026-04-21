package arrays

func Sum_arrays() int {
	array := [5]int{1, 2, 3, 4, 5}
	sum := 0

	for i := range array {
		sum += i
	}

	return sum
}
