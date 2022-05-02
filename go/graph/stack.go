package l337C0d3

type Stack struct {
	size  int
	Items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.size++
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.size == 0 {
		return nil, false
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	s.size--
	return item, true
}

func (s Stack) IsEmpty() bool {
	return s.size == 0
}

func (s Stack) Size() int {
	return s.size
}
