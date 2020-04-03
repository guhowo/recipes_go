package Kahn

import "fmt"

type Graph struct {
	V   int     //顶点个数
	Adj [][]int //adjvex邻接表
}

func (g *Graph) Kahn() {
	//计算入度
	inDegree := make([]int, g.V)
	for i := 0; i < len(g.Adj); i++ {
		for j := 0; j < len(g.Adj[i]); j++ {
			w := g.Adj[i][j]
			inDegree[w]++
		}
	}

	//遍历inDegree，将入度为0的节点放入queue中
	queue := make([]int, 0)
	for _, w := range inDegree {
		if w == 0 {
			queue = append(queue, w)
		}
	}

	//将入度为0的率先取出，并得到剩余的入度为0的节点，继续押入queue
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		fmt.Printf("->%d", top)
		for i := 0; i < len(g.Adj[top]); i++ {
			w := g.Adj[top][i]
			inDegree[w]--
			if inDegree[w] == 0 {
				queue = append(queue, w)
			}
		}
	}
}
