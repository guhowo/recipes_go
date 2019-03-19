package main

import(
	"fmt"
	"sync"
)
func main() {
	var once sync.Once
	i := 0
	f := func(i int) func() {
		return func() {
			fmt.Println("once, index  = ", i)
		}
	}(i)

	done := make(chan struct{}, 1)
	for i:=0; i<10; i++{
		go func(){
			once.Do(f)
			done <- struct{}{}
		}()
	}

	for i:=0; i<10; i++ {
		<- done
	}
}

