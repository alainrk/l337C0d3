package l337C0d3

import (
	"container/heap"
	"fmt"
)

type Point struct {
	x, y, height int
}

type MinHeap []Point // min-heap

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func initBounds(bounds *MinHeap, heightMap [][]int) {
	for i, row := range heightMap {
		for j, height := range row {
			if i == 0 || i == len(heightMap)-1 || j == 0 || j == len(heightMap[0])-1 {
				heap.Push(bounds, Point{x: i, y: j, height: height})
			}
		}
	}
}

func getNeighbours(p Point, lastx int, lasty int) [][]int {
	res := [][]int{}

	directions := [][]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	for _, d := range directions {
		x := p.x + d[0]
		y := p.y + d[1]
		if x < 0 || x > lastx || y < 0 || y > lasty {
			continue
		}
		res = append(res, []int{x, y})
	}
	return res
}

func trapRainWater(heightMap [][]int) int {
	res := 0
	W, H := len(heightMap), len(heightMap[0])
	bounds := &MinHeap{}

	visited := make([][]bool, W)
	for i := range visited {
		visited[i] = make([]bool, H)
	}

	heap.Init(bounds)
	initBounds(bounds, heightMap)

	maxh := -1

	for bounds.Len() > 0 {
		curr := heap.Pop(bounds).(Point) // get min height in the current bounds
		fmt.Println("Pop:", curr)
		visited[curr.x][curr.y] = true
		maxh = max(maxh, curr.height)
		fmt.Println("Maxh:", maxh)
		neighbours := getNeighbours(curr, W-1, H-1)
		for _, nb := range neighbours {
			n := Point{nb[0], nb[1], heightMap[nb[0]][nb[1]]}
			fmt.Println("Neigh:", n)
			if visited[n.x][n.y] {
				continue
			}
			visited[n.x][n.y] = true
			fmt.Printf("n.height (%d) < maxh (%d): %v\n", n.height, maxh, n.height < maxh)
			if n.height < maxh {
				res += (maxh - n.height)
			}
			fmt.Println("res:", res)
			heap.Push(bounds, n)
		}
	}

	fmt.Printf("PQ: %+v\n", bounds)
	fmt.Printf("Visited: %+v\n", visited)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
