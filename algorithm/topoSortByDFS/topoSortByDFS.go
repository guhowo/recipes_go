package topoSortByDFS

import "fmt"

type ISort interface {
	Sort()
}

type Graph struct {
	V   int
	Adj [][]int //邻接表
}

//O(V+E)（V 表示顶点个数，E 表示边的个数）
func (g *Graph) Sort() {
	//通过邻接表生成逆邻接表
	inverseAdj := make([][]int, g.V)
	for i := 0; i < len(g.Adj); i++ {
		for j := 0; j < len(g.Adj[i]); j++ {
			w := g.Adj[i][j]
			inverseAdj[w] = append(inverseAdj[w], i)
		}
	}

	visited := make([]bool, g.V)
	//深度优先遍历图
	for i := 0; i < g.V; i++ {
		if !visited[i] {
			visited[i] = true
			dfs(i, inverseAdj, visited)
		}
	}
}

func dfs(vertex int, inverseAdj [][]int, visited []bool) {
	for i := 0; i < len(inverseAdj[vertex]); i++ {
		w := inverseAdj[vertex][i] //w表示vertex依赖的节点，因此w要先行
		if visited[w] == true {
			continue
		}
		visited[w] = true
		dfs(w, inverseAdj, visited)
	}
	fmt.Printf("--> %d ", vertex)
}
