package main

import (
	"sync"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Service struct {
	ch chan bool
	wg *sync.WaitGroup
}

func NewService() *Service {
	s := &Service {
		ch: make(chan bool),
		wg: &sync.WaitGroup{},
	}

	return s
}

func (s *Service) Stop() {
	close(s.ch)
	s.wg.Wait()
}

func (s *Service) Serve() {
	s.wg.Add(1)
	defer s.wg.Done()

	for {
		select {
		case <- s.ch:
			fmt.Println("main stopping...")
			return
		default:
		}
		s.wg.Add(1)
		go s.anotherServer()
	}
}

func (s *Service) anotherServer() {
	defer s.wg.Done()
	for {
		select {
		case <-s.ch:
			fmt.Println("another stopping...")
			return
		default:
		}
		time.Sleep(time.Second*1)
		//do something
	}
}
func main() {
	service := NewService()
	go service.Serve()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)

	service.Stop()
}