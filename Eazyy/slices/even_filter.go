package slices

func Even_filter(arr []int) []int {
	new_slice := []int{}
	for _, i := range arr {
		if i%2 == 0 {
			new_slice = append(new_slice, i)
		}
	}
	return new_slice
}
