package l337C0d3

import "fmt"

func validParenthesis(str string) bool {
	stack := []rune{}
	pars := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, p := range []rune(str) {
		_, ok := pars[p]

		if ok {
			stack = append(stack, p)
			continue
		}

		if len(stack) == 0 {
			return false
		}
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pars[r] != p {
			return false
		} else {
			fmt.Println("Not recognized:", p)
		}

	}
	return len(stack) == 0
}
