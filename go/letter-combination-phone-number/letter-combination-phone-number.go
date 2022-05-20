package l337C0d3

var keys = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func lcHelper(digits string, cache map[string][]string, solution map[string]bool) []string {
	if v, ok := cache[digits]; ok {
		return v
	}
	res := []string{}
	for i := 0; i < len(digits)-1; i++ {
		r := rune(digits[i])
		_, ok := keys[r]
		if !ok {
			panic("no digit found")
		}
		rest := lcHelper(digits[i+1:], cache, solution)
		for _, d := range keys[r] {
			for _, s := range rest {
				sol := string(d) + s
				if _, ok := solution[sol]; !ok {
					res = append(res, sol)
					solution[sol] = true
				}
			}
		}
	}
	cache[digits] = res
	return res
}

func letterCombinations(digits string) []string {
	cache := make(map[string][]string)
	solution := make(map[string]bool)

	for d, runes := range keys {
		cache[string(d)] = make([]string, len(runes))
		for i, r := range runes {
			cache[string(d)][i] = string(r)
		}
	}

	res := lcHelper(digits, cache, solution)
	return res
}
