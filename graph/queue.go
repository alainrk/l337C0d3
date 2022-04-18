package l337C0d3

type Queue struct {
	size  int
	Items []interface{}
}

func (s *Queue) Enqueue(item interface{}) {
	s.size++
	s.Items = append(s.Items, item)
}

func (s *Queue) Dequeue() (interface{}, bool) {
	if s.size == 0 {
		return nil, false
	}
	item := s.Items[0]
	s.Items = s.Items[1:]
	s.size--
	return item, true
}

func (s Queue) IsEmpty() bool {
	return s.size == 0
}
