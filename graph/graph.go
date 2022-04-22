package l337C0d3

type GraphNode struct {
	Value interface{}
}

type EdgesMap map[int][]int

type Graph struct {
	Nodes []GraphNode
	Edges EdgesMap
}

func (g *Graph) DFS() []interface{} {
	var visit []interface{}
	visited := make([]bool, len(g.Nodes))

	stack := Stack{}

	for {
		// Avoid missing subgraph in not connected graphs
		if stack.IsEmpty() {
			for i, _ := range g.Nodes {
				if !visited[i] {
					stack.Push(i)
					break
				}
			}
		}

		// No more nodes to visit
		if stack.IsEmpty() {
			break
		}

		for !stack.IsEmpty() {
			// Cannot be empty at this point
			i, _ := stack.Pop()
			nodeIdx := i.(int)
			node := g.Nodes[nodeIdx]

			// Visit myself, if not done already
			if !visited[nodeIdx] {
				visited[nodeIdx] = true
				visit = append(visit, node.Value)
			}

			for _, neighbourIdx := range g.Edges[nodeIdx] {
				if !visited[neighbourIdx] {
					stack.Push(neighbourIdx)
				}
			}
		}
	}

	return visit
}

func (g *Graph) BFS() []interface{} {
	var visit []interface{}
	visited := make(map[int]bool, len(g.Nodes))

	// queue := NewQueue(0)
	queue := NewQueue()

	for {
		// Avoid missing subgraph in not connected graphs
		if queue.IsEmpty() {
			for i, _ := range g.Nodes {
				if !visited[i] {
					queue.Enqueue(i)
					break
				}
			}
		}

		// No more nodes to visit
		if queue.IsEmpty() {
			break
		}

		for !queue.IsEmpty() {
			i, _ := queue.Dequeue()
			nodeIdx := i.(int)
			node := g.Nodes[nodeIdx]

			if !visited[nodeIdx] {
				visited[nodeIdx] = true
				visit = append(visit, node.Value)
			}

			for _, neighbourIdx := range g.Edges[nodeIdx] {
				if !visited[neighbourIdx] {
					queue.Enqueue(neighbourIdx)
				}
			}
		}
	}

	return visit
}
