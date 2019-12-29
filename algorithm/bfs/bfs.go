package bfs

import "fmt"

type Graph struct {
	V   int     //顶点个数
	Adj [][]int //adjacent matrix
}

//从图的s到t的路径
func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}

	queue := make([]int, 0)
	visited := make([]bool, g.V)
	prev := make([]int, g.V) //记录前驱
	for i, _ := range prev {
		prev[i] = -1
	}

	queue = append(queue, s)
	visited[s] = true

	poll := 0
	for len(queue) != 0 {
		poll = queue[0]
		queue = queue[1:]
		for _, w := range g.Adj[poll] {
			if !visited[w] {
				prev[w] = poll
				if w == t {
					print(prev, s, t)
					return
				}
				visited[w] = true
				queue = append(queue, w)
			}
		}
	}
}

func print(prev []int, s, t int) {
	if prev[t] != -1 && t != s {
		print(prev, s, prev[t])
	}
	fmt.Println(t, " ")
}
