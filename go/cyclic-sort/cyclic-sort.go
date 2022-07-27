package l337C0d3

func cyclicSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for arr[i] != i {
			arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
		}
	}
	return arr
}
