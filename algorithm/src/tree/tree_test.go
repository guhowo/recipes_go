package tree

import (
	"testing"
)

func TestNode_Traversal1(t *testing.T) {

	a := []int{10, 5, 14, 1, 2, 12, 15, 3}

	var root *Node
	for _, v := range a {
		root = Insert(v, root)
	}

	Traversal1(root)

	for _, unit := range []struct{
		r *Node
		excepted int
	} {
		{root,
		5,
		},
	} {
		if actually := Depth(unit.r); actually != unit.excepted {
			t.Errorf("Depth: [%v], actually: [%v]", unit, actually)
		}
	}

}
