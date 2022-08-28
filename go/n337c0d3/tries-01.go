package l337C0d3

type Trie struct {
	char     byte
	children map[byte]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		char:     'a',
		children: map[byte]*Trie{},
		isEnd:    false,
	}
}

func (this *Trie) Insert(word string) {
	curr := this
	for i, char := range []byte(word) {
		isLast := i == len(word)-1

		if _, ok := curr.children[char]; ok {
			if isLast {
				curr.children[char].isEnd = true
				return
			}
		} else {
			curr.children[char] = &Trie{
				char:     char,
				children: map[byte]*Trie{},
				isEnd:    isLast,
			}
		}
		if isLast {
			return
		}
		curr = curr.children[char]
	}
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, char := range []byte(word) {
		if _, ok := curr.children[char]; !ok {
			return false
		}
		curr = curr.children[char]
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, char := range []byte(prefix) {
		if _, ok := curr.children[char]; !ok {
			return false
		}
		curr = curr.children[char]
	}
	return true
}
