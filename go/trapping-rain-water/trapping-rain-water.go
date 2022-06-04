package l337C0d3

import "fmt"

func trap(height []int) int {
	maxh := 0 // last barrier height
	water := make([]int, len(height))
	res := 0

	for i, h := range height {
		w := maxh - h // max possible level of the water at this point - current level
		if w > 0 {
			water[i] = w
		}
		maxh = max(maxh, h)
	}

	fmt.Println("1", water)

	maxh = 0
	for i := len(water) - 1; i > -1; i-- {
		h := height[i]
		w := min(maxh-h, water[i])
		if w > 0 {
			water[i] = w
			res += w
		}
		maxh = max(maxh, h)
	}

	fmt.Println("2", water)
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
