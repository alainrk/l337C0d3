package l337C0d3

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	nodes := []GraphNode{
		{Value: "0"},
		{Value: "1"},
		{Value: "2"},
		{Value: "3"},
		{Value: "4"},
		{Value: "5"},
		{Value: "6"},
		{Value: "7"},
	}
	edges := EdgesMap{
		0: []int{1, 2},
		1: []int{0, 2, 3, 5},
		2: []int{1},
		3: []int{1},
		4: []int{5},
		5: []int{1, 4},
		6: []int{7},
		7: []int{6},
	}
	graph := Graph{
		Nodes: nodes,
		Edges: edges,
	}

	fmt.Printf("DFS: %+v\n", graph.DFS())
}
