package dynamicArray

import "fmt"

type DynamicArray struct {
	Array    []int
	Length   int
	Capacity int
}

func CreateArray(capacity int) *DynamicArray {
	return &DynamicArray{
		Array:    make([]int, capacity),
		Length:   0,
		Capacity: capacity,
	}
}

func (o *DynamicArray) Put(element int) {
	if o.Length == o.Capacity {
		o.resize()
	}
	o.Array[o.Length] = element
	o.Length++
}

func (o *DynamicArray) resize() {
	newCapacity := o.Capacity * 2

	copyArray := make([]int, newCapacity)
	copy(copyArray, o.Array[:o.Capacity])
	o.Array = copyArray
	o.Capacity = newCapacity
}

func (o *DynamicArray) Print() {
	for i := 0; i < o.Length; i++ {
		fmt.Printf("%d ", o.Array[i])
	}
	fmt.Printf("\n")
}
