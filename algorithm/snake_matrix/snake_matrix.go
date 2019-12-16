package snake_matrix

import "fmt"

type Point struct{
	x int	//横坐标
	y int	//纵坐标
	up int	//移动方向 0左上，1右下
}

func New(m, n int) {

	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}



	point := &Point {
		0,
		0,
		0,
	}
	matrix[0][0] = 1


	for i:= 1; i < m*n; i++ {
		x, y := move(point, m, n)
		matrix[x][y] = i+1
	}

	fmt.Println("matrix = ", matrix)
}


func move(p *Point, m, n int) (int, int){
	if p.up == 0 {
		p.x--
		p.y++
	} else {
		p.x++
		p.y--
	}

	if p.x < 0 {
		p.x++
		p.up = 1	//转向右下方
	}
	if p.y < 0 {
		p.y++
		p.up = 0 	//转向左上方
	}
	if p.y >= n {
		p.x += 2
		p.y--
		p.up = 1
	}
	if p.x >= m {
		p.x--
		p.y += 2
		p.up = 0
	}
	return p.x, p.y
}

/*
func GetValue(x, y int) (value int) {

}
*/