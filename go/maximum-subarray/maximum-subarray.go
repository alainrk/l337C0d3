package l337C0d3

func intMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	currMax := -9999999 // TODO do sth better here
	globMax := -9999999 // TODO do sth better here
	for _, v := range nums {
		currMax = intMax(currMax+v, v)
		globMax = intMax(currMax, globMax)
		// fmt.Printf("i = %d, v = %d, cm = %d, gm = %d\n", i, v, currMax, globMax)
	}
	return globMax
}
