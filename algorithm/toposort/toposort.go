/*
 * topo排序--Kahn算法
 */

package toposort

import "fmt"

type Graph struct {
	V   int     //顶点个数
	Adj [][]int //adjvex邻接表
}

func (g *Graph) Kahn() {
	//计算入度inDegree
	inDegree := make([]int, g.V)
	for _, v := range g.Adj {
		for _, w := range v {
			inDegree[w]++
		}
	}

	list := make([]int, 0)
	for i := 0; i < g.V; i++ {
		if inDegree[i] == 0 {
			list = append(list, i)
		}
	}

	for len(list) > 0 {
		v := list[0]
		fmt.Println(v)
		for _, depend := range g.Adj[v] {
			inDegree[depend]--
			if inDegree[depend] == 0 {
				list = append(list, depend)
			}
		}
		list = list[1:]
	}
}
