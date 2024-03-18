package mspool

import (
	"errors"
	"sync"
	"time"

	// "golang.org/x/text"
)

type sig struct{}

var (
	ErrorInvalidCap = errors.New("cap can not <= 0")
	ErrorInvalidExpire = errors.New("expire can not <= 0")
)

const DefaultExpire = 3
type Pool struct {
	// 容量max
	cap int32
	// running 正在运行
	running int32
	//空闲worker
	sworkers []*Worker
	//expire 过期时间，空闲worker超过这世间回收
	expire time.Duration
	// 释放资源
	release chan sig
	// 保护pool资源安全
	lock sync.Mutex
	// 释放只能调用一次
	once sync.Once
}

func NewPool(cap int) (*Pool, error){
	return NewTimePool(cap, DefaultExpire)
}

func NewTimePool(cap int, expire int) (*Pool, error){
	if cap <=0 {
		return nil, ErrorInvalidCap
	}
	if expire <= 0 {
		return nil, ErrorInvalidExpire
	}
	p := &Pool{
		cap: int32(cap),
		expire: time.Duration(expire)*time.Second,
		release: make(chan sig, 1),
	}
	return p, nil
}

// 提交任务
func (p *Pool) Submit(task func())error{
	// 获取池子一个worker
	w := p.GetWorker()
	w.task <- task
	return nil
}

func (p *Pool)GetWorker() *Worker{
	return nil
}
