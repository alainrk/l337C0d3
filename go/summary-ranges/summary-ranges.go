package l337C0d3

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	res := []string{}

	if len(nums) == 0 {
		return res
	}

	start, end := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		if curr != end+1 {
			var s string
			if start == end {
				s = strconv.Itoa(start)
			} else {
				s = fmt.Sprintf("%d->%d", start, end)
			}
			res = append(res, s)
			start, end = curr, curr
		} else {
			end = curr
		}
	}
	// improve duplication here
	var s string
	if start == end {
		s = strconv.Itoa(start)
	} else {
		s = fmt.Sprintf("%d->%d", start, end)
	}
	res = append(res, s)
	return res
}
