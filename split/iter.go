func iteration(x int) int {
	if x == 1 {
		return 1
	}
	return iteration(x-1) * 2
}
