package dynamicArray

import "testing"

func TestDynamicArray_Put(t *testing.T) {
	array := CreateArray(3)

	for _, unit := range []struct {
		Element int
	}{
		{
			1,
		},
		{
			2,
		},
		{
			3,
		},
		{
			4,
		},
		{
			5,
		},
	} {
		array.Put(unit.Element)
		array.Print()
	}
}
