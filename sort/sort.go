package main

import (
	"fmt"
	"sort"
)

type StudentScore struct {
	name string
	score int
}

type StudentScores []StudentScore

func (s StudentScores)Len() int {
	return len(s)
}

func (s StudentScores)Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s StudentScores)Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main(){
	students := StudentScores{
		{"sharonpan", 100},
		{"haogu", 110},
		{"shawn", 80},
		{"yoki", 90},
		{"vicky", 105},
	}
	fmt.Println("Default: \n\t", students)
	sort.Sort(students)
	fmt.Println("Is Sorted:\n\t", sort.IsSorted(students))
	fmt.Println("Sorted:\n\t", students)
}
