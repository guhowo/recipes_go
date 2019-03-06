package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	var wg sync.WaitGroup
	tb := &TokenBucket{}
	tb.Set(3, 3)	//速率是3， 桶容量是3

	for i:=0; i<10; i++ {
		wg.Add(1)
		fmt.Println("Create req", i, time.Now())
		go func(i int) {
			if tb.Allow() {
				fmt.Println("Response req ", i , time.Now())
			}
			wg.Done()
		}(i)
		time.Sleep(200*time.Millisecond)
	}
	wg.Wait()
}
