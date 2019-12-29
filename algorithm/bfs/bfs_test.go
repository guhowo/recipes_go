package bfs

import "testing"

func TestGraph_BFS(t *testing.T) {
	g := &Graph{
		V:   8,
		Adj: [][]int{{1, 3}, {0, 4, 2}, {1, 5}, {0, 4}, {1, 3, 5, 6}, {2, 4, 7}, {4, 7}, {5, 6}},
	}
	g.BFS(0, 6)
}
