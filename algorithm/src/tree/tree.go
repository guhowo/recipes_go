package tree

import (
	"fmt"
)

type Node struct {
	data  int
	right *Node
	left  *Node
}

//非递归插入
func Insert(e int, root *Node) *Node {
	node := &Node{
		e,
		nil,
		nil,
	}

	ptr := root
	parent := ptr
	pos := ""

	for ptr != nil {
		parent = ptr
		if ptr.data > e {
			ptr = ptr.left
			pos = "left"
		} else if ptr.data < e {
			ptr = ptr.right
			pos = "right"
		} else {
			return root
		}
	}

	if pos == "left" {
		parent.left = node
	} else if pos == "right" {
		parent.right = node
	} else {
		root = node
	}
	return root
}

//递归版本
func Traverse1(root *Node) {
	if root == nil {
		return
	}

	left := root.left
	right := root.right
	Traverse1(left)
	fmt.Println(root.data)
	Traverse1(right)
}



//递归版本
func Depth(root *Node) int {
	if root == nil {
		return 0
	}

	lDepth := Depth(root.left)
	rDepth := Depth(root.right)
	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}

//非递归版本
func Depth1(root *Node) int {
	if root == nil {
		return 0
	}

	slice := make([]*Node, 0)
	slice = append(slice, root)

	var dpt int

	for len(slice) != 0 {
		dpt++
		cnt := 0
		curLevelTotal := len(slice)
		for cnt < curLevelTotal {
			cnt++
			end := slice[len(slice) - 1]
			slice =slice[:len(slice) - 1]
			if end.right != nil {
				slice = append(slice, end.right)
			}
			if end.left != nil {
				slice = append(slice, end.left)
			}
		}
	}
	return dpt
}