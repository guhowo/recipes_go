package queen8

import (
	"fmt"
)

type Queen interface {
	Queen8(row int, result []int)
	Print(row []int)
}

var total int
var output [][]string

func Queen8(row int, result []int) {
	if row == 8 {
		PrintV2(result)
		total++
		return
	}
	for i := 0; i < 8; i++ {
		if isOk(row, i, result) {
			result[row] = i
			Queen8(row+1, result)
		}
	}
}

func isOk(row, column int, result []int) bool {
	leftup, rightup := column-1, column+1
	for i := row - 1; i >= 0; i-- {
		//在同一列
		if result[i] == column {
			return false
		}
		if leftup >= 0 {
			if result[i] == leftup {
				return false
			}
			leftup--
		}
		if rightup < 8 {
			if result[i] == rightup {
				return false
			}
			rightup++
		}
	}
	return true
}

func Print(result []int) {
	fmt.Println("result = ")
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if result[j] == i {
				fmt.Printf("Q ")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Printf("\n")
	}
}

func PrintV2(result []int) {
	out := make([]string, 8)
	for i := 0; i < 8; i++ {
		rowStr := ""
		for j := 0; j < 8; j++ {
			if result[j] == i {
				//fmt.Printf("Q ")
				rowStr += "Q "
			} else {
				//fmt.Printf("* ")
				rowStr += "* "
			}
		}
		out[i] = rowStr
	}
	output = append(output, out)
}
