package dfs

import "fmt"

type Graph struct {
	V   int     //顶点个数
	Adj [][]int //adjacent matrix
}

var found = false

//从图的s到t的路径
func (g *Graph) DFS(s, t int) {
	if s == t {
		return
	}

	visited := make([]bool, g.V)
	prev := make([]int, g.V) //记录前驱
	for i, _ := range prev {
		prev[i] = -1
	}

	//found := false
	g.recurDFS(s, t, visited, prev)
	print(prev, s, t)
}

func (g *Graph) recurDFS(w, t int, visited []bool, prev []int) {
	if found {
		return
	}
	visited[w] = true
	if w == t {
		found = true
		return
	}
	for _, v := range g.Adj[w] {
		if !visited[v] {
			prev[v] = w
			g.recurDFS(v, t, visited, prev)
		}
	}
}

func print(prev []int, s, t int) {
	if prev[t] != -1 && t != s {
		print(prev, s, prev[t])
	}
	fmt.Println(t, " ")
}
