package l337C0d3

import (
	"fmt"
)

func intMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	currMax := -9999999 // TODO do sth better here
	globMax := -9999999 // TODO do sth better here
	for i, v := range nums {
		currMax = intMax(currMax+v, v)
		globMax = intMax(currMax, globMax)
		fmt.Printf("i = %d, v = %d, cm = %d, gm = %d\n", i, v, currMax, globMax)
	}
	return globMax
}

/*
-2  1 -3  4  -1  2  1  -5  4
-2 -1 -4  0  -1  1  2  -3  1
                           4
                       -1  4
*/
