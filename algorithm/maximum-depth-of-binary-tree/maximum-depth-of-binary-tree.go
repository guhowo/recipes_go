package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归方法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1

	if l > r {
		return l
	}
	return r
}

//BFS
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	deque := make([]*TreeNode, 0)
	deque = append(deque, root)
	level := 0

	for len(deque) > 0 {
		level++
		size := len(deque)
		for size > 0 {
			head := deque[0]
			if head.Left != nil {
				deque = append(deque, head.Left)
			}
			if head.Right != nil {
				deque = append(deque, head.Right)
			}
			deque = deque[1:]
			size--
		}
	}

	return level
}

/*
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	depth := 0
	levelSize := 0
	for len(queue) != 0 {
		depth++
		levelSize = len(queue)
		for i := 0; i < levelSize; i++ {
			p := queue[0]
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
			queue = queue[1:]
		}
	}
	return depth
}
*/
