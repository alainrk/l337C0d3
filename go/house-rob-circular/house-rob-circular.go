package l337C0d3

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solveRobCircular(nums []int, start, end int) int {
	res := [2]int{0, nums[start]}
	for i := start + 1; i < end; i++ {
		res[0], res[1] = res[1], max(res[1], res[0]+nums[i])
	}
	return res[1]
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// Get the max of "excluding the last" and "excluding the first" solutions
	return max(solveRobCircular(nums, 0, len(nums)-1), solveRobCircular(nums, 1, len(nums)))
}
