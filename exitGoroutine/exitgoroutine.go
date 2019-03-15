package main

import (
	"sync"
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func consume(stop <-chan bool, index int)  {
	for {
		select {
		case <-stop:
			fmt.Println("exit goroutine ", index)
			return
		default:
			fmt.Println("goroutine", index, " running...")
			time.Sleep(time.Second*1)
		}
	}
}

//捕获信号，在信号产生前一直阻塞
func waitForSignals() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
}

func main() {
	stop := make(chan bool)
	var wg sync.WaitGroup

	//generate n goroutines
	for i:=0; i<3; i++ {
		wg.Add(1)
		go func(stop <-chan bool, index int) {
			defer wg.Done()
			consume(stop, index)
		}(stop, i)
	}

	waitForSignals()
	close(stop)
	fmt.Println("stopping all jobs!")
	wg.Wait()
}
