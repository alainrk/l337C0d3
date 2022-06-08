package l337C0d3

func missingNumber(nums []int) int {
	n := len(nums)
	sum := (n * (n + 1)) / 2
	for _, x := range nums {
		sum -= x
	}
	return sum
}
