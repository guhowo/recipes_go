package heap

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	a := []float64{-1, 3, 5, 3, 9, 8, 1}
	Sort(a, 1, len(a)-1)
	fmt.Println(a)
}

func TestMaxHeapInsert(t *testing.T) {
	a := []float64{-1, 3, 5, 3, 9, 8, 1}
	BuildHeap(a, 1, len(a)-1)

	fmt.Println("插入前： ", a)
	a = MaxHeapInsert(a, 1, len(a)-1, 7)
	fmt.Println("插入后： ", a)
	a = MaxHeapDelete(a, 1, len(a)-1)
	fmt.Println("删除后： ", a)
	a = MaxHeapDelete(a, 1, len(a)-1)
	fmt.Println("删除后： ", a)
}

func TestMinHeapInsert(t *testing.T) {
	a := []float64{-1, 3, 5, 3, 9, 8, 1}
	BuildHeap(a, 1, len(a)-1)

	fmt.Println("插入前： ", a)
	a = MaxHeapInsert(a, 1, len(a)-1, 7)
	fmt.Println("插入后： ", a)
	a = MaxHeapDelete(a, 1, len(a)-1)
	fmt.Println("删除后： ", a)
	a = MaxHeapDelete(a, 1, len(a)-1)
	fmt.Println("删除后： ", a)
}

func TestMedian(t *testing.T) {
	a := []float64{0, 1, 2, 3, 4, 5, 6}
	m := Median(a, 1, 7)
	fmt.Println(m)
}
