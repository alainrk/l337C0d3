package l337C0d3

type Queue struct {
	size     int
	elements []Element
}

func NewQueue(items ...interface{}) Queue {
	var elements []Element
	for _, e := range items {
		elements = append(elements, Element{Value: e})
	}
	return Queue{
		size:     len(items),
		elements: elements,
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.size++
	q.elements = append(q.elements, Element{Value: item})
}

func (q *Queue) Dequeue() (interface{}, bool) {
	if q.size == 0 {
		return nil, false
	}
	item := q.elements[0]
	q.elements = q.elements[1:]
	q.size--
	return item.Value, true
}

func (q Queue) IsEmpty() bool {
	return q.size == 0
}

func (q Queue) Size() int {
	return q.size
}
