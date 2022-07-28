package l337C0d3

func cyclicSortMissing(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for arr[i] != i {
			curr := arr[i]
			if curr >= len(arr) {
				break
			}
			arr[i], arr[curr] = arr[curr], curr
		}
	}

	if arr[0] != 0 {
		return 0
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1]+1 {
			return i
		}
	}
	return len(arr)
}
