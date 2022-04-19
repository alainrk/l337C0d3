package l337C0d3

import (
	"reflect"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	expected := Queue{
		size:     3,
		elements: []Element{{1}, {2}, {3}},
	}

	if !reflect.DeepEqual(q, expected) {
		t.Errorf("error enqueue, given: %+v expected: %+v", q, expected)
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue(1, 2, 3, 4, 5)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	expected := Queue{
		size:     2,
		elements: []Element{{4}, {5}},
	}

	if !reflect.DeepEqual(q, expected) {
		t.Errorf("error enqueue, given: %+v expected: %+v", q, expected)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		q    Queue
		want bool
	}{
		{name: "Empty Queue", q: NewQueue(), want: true},
		{name: "Non empty Queue", q: NewQueue(1, 2, 3, 4), want: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := Queue{
				size:     tt.q.size,
				elements: tt.q.elements,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("Queue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
