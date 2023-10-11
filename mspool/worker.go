package mspool

import (
	"time"
)

type Worker struct{
	pool *Pool
	task chan func()
	// lasttime最后执行任务时间
	lastTime time.Time
}
