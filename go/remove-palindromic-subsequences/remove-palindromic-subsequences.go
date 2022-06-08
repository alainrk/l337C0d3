package l337C0d3

func removePalindromeSub(s string) int {
	l := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[l-i] {
			return 2
		}
	}
	return 1
}
