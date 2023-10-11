package demotest

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func PPoftmain() {
	// 开启pprof，监听请求
	ip := "127.0.0.1:6069"
	// 开启pprof
	go func() {
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	// 路由，访问，触发内存泄露的代码判断
	http.HandleFunc("/test", handler)

	// 阻塞
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 接收端受到的channel为nil
	ch := make(chan int, 1)
	go func() {
		<-ch
		fmt.Println(111)
	}()
}
