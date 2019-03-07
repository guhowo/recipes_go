package main

import (
	"errors"
	"io"
	"time"
	"sync"
)

var (
	ErrInvalidConfig = errors.New("invalid pool config")
	ErrPoolClosed	 = errors.New("pool closed")
)

type Poolable interface {
	io.Closer
	GetActiveTime() time.Time	//最近活跃时间
}

type Pool interface {
	Acquire() (Poolable, error)		// 获取资源
	Release(Poolable) error			//释放资源
	Close(Poolable)	error			//关闭资源
	Shutdown()		error			//关闭池子
}

type factory func() (Poolable, error)

type GenericPool struct {
	sync.Mutex
	pool chan Poolable
	maxOpen	int
	minOpen int
	numOpen int			// 当前池中资源数
	closed 	bool		// 池是否已关闭
	maxLifeTime	time.Duration
	factory factory		// 创建连接的方法
}

//根据配置创建连接池
func NewGenericPool(minOpen, maxOpen int, maxLifeTime time.Duration, factory factory) (*GenericPool, error) {
	if maxOpen <= 0 || minOpen>maxOpen {
		return nil, ErrInvalidConfig
	}
	p := &GenericPool{
		maxOpen:		maxOpen,
		minOpen:		minOpen,
		maxLifeTime: 	maxLifeTime,
		factory:		factory,
		pool:			make(chan Poolable, maxOpen),
	}
	for i:=0; i<minOpen; i++ {
		closer, err := factory()
		if err != nil {
			continue
		}
		p.numOpen++
		p.pool<-closer
	}
	return p, nil
}

//获取可连接的资源
func (p *GenericPool) Acquire() (Poolable, error) {
	if p.closed {
		return nil, ErrPoolClosed
	}
	for {
		closer, err := p.getOrCreate()
		if err != nil {
			return nil, err
		}
		// 如果设置了超时且当前连接的活跃时间+超时时间早于现在，则当前连接已过期
		if p.maxLifeTime > 0 && closer.GetActiveTime().Add(p.maxLifeTime).Before(time.Now()) {
			closer.Close()
			continue
		}
		return closer, nil
	}
}

func (p *GenericPool) getOrCreate() (Poolable, error) {
	select {
	case closer := <-p.pool:
		return closer, nil
	default:
	}

	p.Lock()
	defer p.Unlock()

	if p.numOpen >= p.maxOpen {
		closer := <-p.pool
		return closer, nil
	}

	closer, err := p.factory()
	if err != nil {
		return nil, err
	}
	p.numOpen++
	return closer, nil
}

// 释放单个资源到连接池
func (p *GenericPool) Release(closer Poolable) error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	defer p.Unlock()
	p.pool <- closer

	return nil
}

//关闭单个资源
func (p *GenericPool) Close(closer Poolable) error {
	p.Lock()
	defer p.Unlock()

	closer.Close()
	p.numOpen--
	return nil
}

// 关闭连接池，释放所有资源
func (p *GenericPool) Shutdown() error {
	if p.closed {
		return ErrPoolClosed
	}

	p.Lock()
	defer p.Unlock()

	for closer := range p.pool {
		closer.Close()
	}
	p.numOpen = 0
	p.closed = true
	return nil
}