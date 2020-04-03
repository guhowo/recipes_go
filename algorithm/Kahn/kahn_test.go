package Kahn

import "testing"

func TestGraph_Kahn(t *testing.T) {
	g := Graph{
		V:   4,
		Adj: [][]int{{1}, {2, 3}, {}, {2}},
	}

	g.Kahn()
}
