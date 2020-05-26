package allrange

import (
	"fmt"
	"testing"
	"time"
)

func TestAllRange(t *testing.T) {
	cur := []int{-1, -1, 2}
	allRange(cur, 0, len(cur))
	fmt.Println(result)
}

func Test_Sharon(t *testing.T) {
	quit := make(chan struct{})
	done := make(chan struct{})
	ch := make(chan int, 100)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	go func() {
	Loop:
		for {
			select {
			case <-quit:
				fmt.Println("quit")
				done <- struct{}{}
				break Loop
			case i, ok := <-ch:
				if !ok {
					fmt.Println("closed")
					done <- struct{}{}
					return
				}
				fmt.Println(i)
				time.Sleep(time.Second * 1)
			}

		}
	}()

	<-done
}

func Test_Sharonv2(t *testing.T) {
	ch := make(chan int, 100)

	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		select {
		case i, ok := <-ch:
			if !ok {
				fmt.Println("closed")
				return
			}
			fmt.Println(i)
			time.Sleep(time.Second * 1)
		}

	}
}
