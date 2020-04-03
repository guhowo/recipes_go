package topoSortByDFS

import "testing"

func TestGraph_Sort(t *testing.T) {
	var isort ISort
	g := &Graph{
		V:   4,
		Adj: [][]int{{1}, {2, 3}, {}, {2}},
	}
	isort = g
	isort.Sort()
}
