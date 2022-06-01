package l337C0d3

import "math"

func maxPoints(points [][]int) int {
	res := 0

	for i := 0; i < len(points); i++ {
		slopes := make(map[float64]int)
		for j := i + 1; j < len(points); j++ {
			s := calcSlope(points[i], points[j])
			if _, ok := slopes[s]; !ok {
				slopes[s] = 0
			}
			slopes[s]++
			res = max(res, slopes[s])
		}
	}

	return res + 1
}

func calcSlope(p1, p2 []int) float64 {
	if p2[0]-p1[0] == 0 {
		return math.Inf(1)
	}
	return float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
