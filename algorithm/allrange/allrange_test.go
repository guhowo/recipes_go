package allrange

import (
	"fmt"
	"testing"
)

func TestAllRange(t *testing.T) {
	cur := []int{-1, -1, 2}
	allRange(cur, 0, len(cur))
	fmt.Println(result)
}
