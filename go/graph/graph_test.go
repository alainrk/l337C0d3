package l337C0d3

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	nodes := []GraphNode{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}}
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

	// One of the possible result, but the expected for the current implementation
	expected := []int{0, 2, 1, 5, 4, 3, 6, 7}
	dfs := graph.DFS()
	res := make([]int, len(dfs))
	for i, v := range dfs {
		res[i] = v.(int)
	}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("DFS error, expected: %+v, given: %+v", expected, res)
	}
}

func TestBFS(t *testing.T) {
	nodes := []GraphNode{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}}
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

	// One of the possible result, but the expected for the current implementation
	expected := []int{0, 1, 2, 3, 5, 4, 6, 7}
	bfs := graph.BFS()
	res := make([]int, len(bfs))
	for i, v := range bfs {
		res[i] = v.(int)
	}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("BFS error, expected: %+v, given: %+v", expected, res)
	}
}

func TestBFS2(t *testing.T) {
	nodes := []GraphNode{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}}
	edges := EdgesMap{
		0: []int{},
		1: []int{},
		2: []int{},
		3: []int{},
		4: []int{},
		5: []int{},
		6: []int{},
		7: []int{},
	}
	graph := Graph{
		Nodes: nodes,
		Edges: edges,
	}

	// One of the possible result, but the expected for the current implementation
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7}
	bfs := graph.BFS()
	res := make([]int, len(bfs))
	for i, v := range bfs {
		res[i] = v.(int)
	}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("BFS error, expected: %+v, given: %+v", expected, res)
	}
}

func TestBFS3(t *testing.T) {
	nodes := []GraphNode{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}}
	edges := EdgesMap{
		0: []int{1},
		1: []int{2},
		2: []int{3},
		3: []int{4},
		4: []int{5},
		5: []int{6},
		6: []int{7},
		7: []int{0},
	}
	graph := Graph{
		Nodes: nodes,
		Edges: edges,
	}

	// One of the possible result, but the expected for the current implementation
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7}
	bfs := graph.BFS()
	res := make([]int, len(bfs))
	for i, v := range bfs {
		res[i] = v.(int)
	}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("BFS error, expected: %+v, given: %+v", expected, res)
	}
}

func TestBFS4(t *testing.T) {
	nodes := []GraphNode{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}}
	edges := EdgesMap{
		0: []int{7},
		1: []int{},
		2: []int{},
		3: []int{},
		4: []int{},
		5: []int{},
		6: []int{},
		7: []int{0, 1, 2, 3, 4, 5, 6},
	}
	graph := Graph{
		Nodes: nodes,
		Edges: edges,
	}

	// One of the possible result, but the expected for the current implementation
	expected := []int{0, 7, 1, 2, 3, 4, 5, 6}
	bfs := graph.BFS()
	res := make([]int, len(bfs))
	for i, v := range bfs {
		res[i] = v.(int)
	}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("BFS error, expected: %+v, given: %+v", expected, res)
	}
}
