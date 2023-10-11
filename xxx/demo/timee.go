package demotest

import (
	"fmt"
	"time"
)

func TickerTest() {
	//定义一个ticker，每隔500毫秒触发
	ticker := time.NewTicker(time.Second*1)
	//Ticker触发
	go func() {
		for t := range ticker.C {
			fmt.Println("ticker被触发", t)
		}
	}()

	time.Sleep(time.Second * 10)
	//停止ticker
	ticker.Stop()
}
