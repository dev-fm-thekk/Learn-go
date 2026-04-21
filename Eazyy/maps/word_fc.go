package maps

func Word_fc(words []string) map[string]int {
	count := make(map[string]int)

	for _, i := range words {
		_, ok := count[i]
		if !ok {
			count[i] = 1
		} else {
			count[i] += 1
		}
	}

	return count
}
