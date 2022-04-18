package l337C0d3

type GraphNode struct {
	Value interface{}
}

type Graph struct {
	Nodes []*GraphNode
	Edges map[int]int
}

func (g *Graph) DFS() []GraphNode {
	return nil
}

func (g *Graph) BFS() []GraphNode {
	return nil
}
