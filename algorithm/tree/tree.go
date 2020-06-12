package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

var root *Node = nil

func insert(val int) {
	node := Node{Val: val}
	if root == nil {
		root = &node
		return
	}

	p := root
	for p != nil {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &node
				return
			}
			p = p.Left
		} else if val > p.Val {
			if p.Right == nil {
				p.Right = &node
				return
			}
			p = p.Right
		} else {
			return
		}
	}

	return
}

func find(val int) *Node {
	p := root

	for p != nil {
		if p.Val == val {
			return p
		} else if p.Val > val {
			p = p.Left
		} else {
			p = p.Right
		}
	}

	return nil
}

func delete(val int) {
	p := root
	var pp *Node = nil //要删除节点的parent

	//先找到要删除的节点和它的parent
	for p != nil {
		if val < p.Val {
			pp = p
			p = p.Left
		} else if val > p.Val {
			pp = p
			p = p.Right
		} else {
			break
		}
	}

	//没有找到要删除的节点，退出
	if p == nil {
		return
	}

	//要删除的节点有两个子节点，找到右子树的最小的节点minp（minp肯定没有叶子节点）
	if p.Left != nil && p.Right != nil {
		minp := p.Right
		var minpp *Node = nil //minp的父亲节点
		for minp != nil {
			if minp.Left != nil {
				minpp = minp
				minp = minp.Left
			}
		}
		p.Val = minp.Val
		p = minp   //要删除的点解变成了minp
		pp = minpp //要删除节点的父节点变成了minpp
	}

	// 找到要删除节点的孩子节点（最多一个，因为两个孩子的情况上一段已经转换成最多一个的情况了）
	var child *Node = nil
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}

	//将要删除节点的父节点与要删除节点的子节点链接起来
	if pp == nil { //要删除的节点是root
		root = child
	} else if pp.Left == p {
		pp.Left = child
	} else {
		pp.Right = child
	}
}

// 前序遍历
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

//前序遍历的深度优先实现
func PreOrder(root *Node) {
	if root == nil {
		return
	}

	stack := make([]*Node, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		//左子树入stack
		for cur != nil {
			fmt.Println(cur.Val)
			stack = append(stack, cur) //push
			cur = cur.Left
		}
		//左子树遍历完毕，遍历右子树
		if len(stack) != 0 {
			cur = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1] //pop
			cur = cur.Right
		}
	}
}

//中序遍历的深度优先
func InOrder(root *Node) {
	if root == nil {
		return
	}

	stack := make([]*Node, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		//左子树入stack
		for cur != nil {
			stack = append(stack, cur) //push
			cur = cur.Left
		}
		//左子树遍历完毕，遍历右子树
		if len(stack) != 0 {
			cur = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1] //pop
			fmt.Println(cur.Val)
			cur = cur.Right
		}
	}
}

// 后序遍历
func PostOrder(root *Node) {
	if root == nil {
		return
	}

	stack := make([]*Node, 0)
	cur := root
	var pre *Node = nil

	for cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			pre = cur
			cur = cur.Left
		}

		if len(stack) != 0 {
			p := stack[len(stack)-1]
			if (p.Left == nil && p.Right == nil) || pre == p.Right {
				fmt.Println(p.Val)
				stack = stack[0 : len(stack)-1]
			} else {
				if p.Right != nil {
					cur = p.Right
				}
			}

		}
	}
}

//按行遍历
func bfs(root *Node) {
	if root == nil {
		return
	}

	deque := make([]*Node, 0)
	deque = append(deque, root)

	for len(deque) != 0 {
		levelSize := len(deque)
		for i := 0; i < levelSize; i++ {
			pop := deque[0]
			fmt.Println(pop.Val)
			deque = deque[1:]
			if pop.Left != nil {
				deque = append(deque, pop.Left)
			}
			if pop.Right != nil {
				deque = append(deque, pop.Right)
			}
		}
	}
}
