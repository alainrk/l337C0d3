package l337C0d3

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

func rob(nums []int) int {
	memo := map[int]int{}
	return solveRob(nums, len(nums)-1, memo)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
