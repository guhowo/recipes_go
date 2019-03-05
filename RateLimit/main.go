package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	lr := &LimitRate{}

	lr.Set(3, time.Second)	// 1s内最多请求3次

	for i:=0; i< 10; i++ {
		wg.Add(1)
		fmt.Println("create req", i , time.Now())
		go func(i int) {
			if lr.Allow2() {
				fmt.Println("Response", i, time.Now())
			}
			wg.Done()
		}(i)
		time.Sleep(200*time.Millisecond)	//每200ms发出一个请求
	}
	wg.Wait()
	fmt.Println("cost ", time.Now().Sub(start))
}