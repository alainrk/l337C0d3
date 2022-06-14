package l337C0d3

import "fmt"

func search(nums []int, target int) int {
	if nums[0] <= nums[len(nums)-1] {
		return bsearch(nums, target)
	}

	newzero := getZero(nums)
	nums = append(nums[newzero:], nums[:newzero]...)

	if res := bsearch(nums, target); res >= 0 {
		return (res + newzero) % len(nums)
	}
	return -1
}

func getZero(nums []int) int {
	s, e := 0, len(nums)-1
	zero := -1

	for s <= e {
		m := (e + s) / 2
		fmt.Println(nums[m], nums[s:e+1])

		// only 2 elements left
		if e-s == 1 {
			if e > 0 && nums[e] < nums[s] {
				zero = e
				break
			}
			zero = s
			break
		}

		if nums[s] > nums[m] {
			e = m
			continue
		}

		s = m
	}
	return zero
}

func bsearch(nums []int, target int) int {
	s, e := 0, len(nums)
	for s < e {
		mid := (s + e) / 2
		result := nums[mid] - target
		if result < 0 {
			s = mid + 1
		} else if result == 0 {
			return mid
		} else {
			e = mid
		}
	}
	return -1
}

/*
Split thing

7 8 1 2 3 4 5 6
7 8 1 2
8 1 2
8 1

i | i > 1 AND a[i] < a[i-1] => i IS a[0]
*/
