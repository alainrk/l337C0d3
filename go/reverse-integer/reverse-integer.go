package l337C0d3

import "strconv"

func pad(s string, n int) string {
	for i := 0; i < n-len(s); i++ {
		s = "0" + s
	}
	return s
}

func isValid(s string, neg bool) bool {
	if len(s) > 10 {
		return false
	}
	if neg && s > "2147483648" {
		return false
	}
	if !neg && s > "2147483647" {
		return false
	}
	return true
}

func reverse(x int) int {
	neg := x < 0
	s := []byte(strconv.Itoa(x))
	if neg {
		s = s[1:]
	}
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}

	sb := pad(string(s), 10)

	if !isValid(sb, neg) {
		return 0
	}

	i, _ := strconv.Atoi(sb)
	if neg {
		i *= -1
	}
	return i
}
