package main

import (
	"log"
	"sync"
	"math/rand"
	"time"
)

//一个生产者多个消费者
func main() {
	log.SetFlags(0)
	rand.Seed(time.Now().UnixNano())
	const (
		MaxRandomNumber = 100000
		NumReceivers	= 100
		)

	var wg sync.WaitGroup
	ch := make(chan int, 100)

	//一个生产者
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0{
				close(ch)
				return
			} else {
				ch <- value
			}
		}
	}()

	//多个消费者
	for i:=0; i< NumReceivers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range ch {
				log.Println(value)
			}
		}()
	}

	wg.Wait()
}