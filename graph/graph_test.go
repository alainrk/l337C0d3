package l337C0d3

import (
	"reflect"
	"testing"
)

func TestGraph_DFS(t *testing.T) {
	type fields struct {
		Nodes []*GraphNode
		Edges map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   []GraphNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := &Graph{
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			if got := g.DFS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_BFS(t *testing.T) {
	type fields struct {
		Nodes []*GraphNode
		Edges map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   []GraphNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := &Graph{
				Nodes: tt.fields.Nodes,
				Edges: tt.fields.Edges,
			}
			if got := g.BFS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
