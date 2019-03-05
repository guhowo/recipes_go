package main

import (
	"time"
	"sync"
)


type LimitRate struct {
	rate 	int					//计数周期内最多允许的请求数
	begin 	time.Time			//计数开始时间
	cycle 	time.Duration		//计数周期
	count 	int					//计数周期内累计收到的请求数
	lock 	sync.Mutex
}


//version 1: ignore the request when overload
func (l *LimitRate) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.count >= l.rate - 1 {
		now :=  time.Now()
		if now.Sub(l.begin) >= l.cycle {
			l.Reset(now)
			return true
		} else {
			return false
		}
	} else {
		l.count++
		return true
	}
}

//version 2: loop and wait when overload
func (l *LimitRate) Allow2() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.count >= l.rate - 1 {
		for {
			now := time.Now()
			if now.Sub(l.begin) >= l.cycle {
				l.Reset(now)
				return true
			} else {
				//do nothing to wait
			}
		}
	} else {
		l.count++
		return true
	}
}

//set执行的前提是已经取得限速器的锁，所以这里不加锁
func (l *LimitRate) Set(r int, cycle time.Duration) {
	l.rate 	= r
	l.begin = time.Now()
	l.cycle = cycle
	l.count = 0
}

//同set
func (l *LimitRate) Reset(t time.Time) {
	l.begin = t
	l.count = 0
}