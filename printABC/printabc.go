package main

import (
	"fmt"
	"time"
)

func main() {
	chA := make(chan struct{})
	chB := make(chan struct{})
	chC := make(chan struct{})

	start := 0
	//start
	go func() {
		chA <- struct{}{}
	}()

	go func() {
		for{
			<- chA
			fmt.Println("1st gorounte: ", start)
			start++
			time.Sleep(time.Second*1)
			chB <- struct{}{}
		}
	}()

	go func() {
		for{
			<- chB
			fmt.Println("2st gorounte: ", start)
			start++
			time.Sleep(time.Second*1)
			chC <- struct{}{}
		}
	}()

	go func() {
		for{
			<- chC
			fmt.Println("3st gorounte: ", start)
			start++
			time.Sleep(time.Second*1)
			chA <- struct{}{}
		}
	}()

	select {}
}
