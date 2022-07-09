package l337C0d3

func wordBreak(s string, d []string) bool {
	dict := map[string]bool{}
	for _, x := range d {
		dict[x] = true
	}
	return wb(s, dict, map[string]bool{})
}

func wb(s string, dict map[string]bool, memo map[string]bool) bool {
	if _, ok := dict[s]; ok || s == "" {
		return true
	}

	if v, ok := memo[s]; ok {
		return v
	}

	for i := 1; i <= len(s); i++ {
		w := s[:i]

		if _, ok := dict[w]; ok {
			if wb(s[i:], dict, memo) {
				memo[s] = true
				return true
			}
		}
	}

	memo[s] = false
	return false
}
