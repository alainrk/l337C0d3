package l337C0d3

type pos [2]int

func solve(board [][]byte) [][]byte {
	m, n := len(board), len(board[0])
	queue := []pos{} // can be improved of course

	// set non-border O as T, so that I can detect which is captured if not changed back to O
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}

			// on the border => keep it as O and push in the queue
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				queue = append(queue, pos{i, j})
				continue
			}
			// not on the border => set temporarily as T
			board[i][j] = 'T'
		}
	}

	for len(queue) > 0 {
		// pop
		curr := queue[0]
		queue = queue[1:]

		for _, neigh := range getNeighbours(curr, m, n) {
			if board[neigh[0]][neigh[1]] == 'X' || board[neigh[0]][neigh[1]] == 'O' {
				continue
			}
			// Mark "visited" and non-captured and push it in the queue
			board[neigh[0]][neigh[1]] = 'O'

			queue = append(queue, neigh)
		}
	}

	// replace all the captured 'T' to 'X'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'T' {
				board[i][j] = 'X'
			}
		}
	}

	return board
}

func getNeighbours(p pos, m, n int) []pos {
	res := make([]pos, 4)
	directions := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	for _, d := range directions {
		i, j := p[0]+d[0], p[1]+d[1]
		if i < 0 || i >= m || j < 0 || j >= n {
			continue
		}
		res = append(res, pos{i, j})
	}
	return res
}
