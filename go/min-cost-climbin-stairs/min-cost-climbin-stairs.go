package l337C0d3

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solveMinCostCS(cost []int, step int, memo map[int]int) int {
	if step < 1 {
		return 0
	}
	if v, ok := memo[step]; ok {
		return v
	}
	memo[step] = cost[step] + min(solveMinCostCS(cost, step-1, memo), solveMinCostCS(cost, step-2, memo))
	return memo[step]
}

func minCostClimbingStairs(cost []int) int {
	memo := map[int]int{}
	cost = append(cost, 0)
	cost = append([]int{0}, cost...)
	return solveMinCostCS(cost, len(cost)-1, memo)
}

/*
  MinCost(n) = Min(MinCost(n-1), MinCost(n-2)) + Cost(n)
*/
