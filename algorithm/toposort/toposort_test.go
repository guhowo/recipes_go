package toposort

import (
	"testing"
)

func Test_TopoSortByKahn(t *testing.T) {
	g := &Graph{
		V:   4,
		Adj: [][]int{{1}, {2, 3}, {}, {2}},
	}
	g.Kahn()
}
