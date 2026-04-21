package loops

func Sum_n(n int) int {
	var acc int = 0

	for i := 1; i <= n; i++ {
		acc += i
	}

	return acc
}
