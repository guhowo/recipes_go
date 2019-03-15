package main

import (
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
	"fmt"
)

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGHUP)
	ctx, cancel := context.WithCancel(context.Background())
	line := make(chan int)
	exit := make(chan struct{})

	go Producer(line)
	go Work(ctx, line, exit)
	<-sig
	cancel()
	<-exit
}

func Producer(line chan int) {
	for i:=0; i<10; i++ {
		line <- i
		time.Sleep(time.Second*1)
	}
}

func Work(ctx context.Context, line chan int, exit chan struct{}) {
	for {
		select {
		case n := <- line:
			fmt.Println("goroutine running...", n)
		case <-ctx.Done():
			fmt.Println("exiting")
			goto EXIT
		}
	}
	EXIT:
		exit<- struct{}{}
}