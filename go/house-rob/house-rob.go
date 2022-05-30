package l337C0d3

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solveRob(nums []int, pos int, memo map[int]int) int {
	if pos < 0 {
		return 0
	}
	if v, ok := memo[pos]; ok {
		return v
	}
	memo[pos] = max(solveRob(nums, pos-1, memo), solveRob(nums, pos-2, memo)+nums[pos])
	return memo[pos]
}

// Top down
func rob(nums []int) int {
	memo := map[int]int{}
	return solveRob(nums, len(nums)-1, memo)
}

// Bottom up
func robIter(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := [2]int{0, nums[0]}
	for i := 1; i < len(nums); i++ {
		res[0], res[1] = res[1], max(res[1], res[0]+nums[i])
	}
	return res[1]
}
