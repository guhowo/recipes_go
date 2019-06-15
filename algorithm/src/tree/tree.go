package tree

import "fmt"

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


func Traversal1(root *Node) {
	if root == nil {
		return
	}

	left := root.left
	right := root.right
	Traversal1(left)
	fmt.Println(root.data)
	Traversal1(right)
}

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