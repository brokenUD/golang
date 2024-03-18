package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func Func1(ctx context.Context) {
	ctx = context.WithValue(ctx, "k1", "v1")
	func2(ctx)
}

func func2(ctx context.Context) {
	fmt.Println(ctx.Value("k1").(string))
}

// timeout := 10 * time.Second
// t = time.Now().Add(timeout)
// conn.SetDeadline(t)

func Func2() error {
	respC := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		respC <- 10
		close(respC)
	}()

	select {
	case r := <-respC:
		fmt.Printf("Resp: %d\n", r)
		return nil
	case <-time.After(time.Second * 2):
		fmt.Println("catch timeout")
		return errors.New("timeout")
	}
}


func func3(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	respC := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		respC <- 10
	}()

	// sl := &sync.Mutex{}
	// sl.Lock()
	// sl.Unlock()
	// sl.Lock()
	// sl.Unlock()
	select {
	case <-ctx.Done():
		fmt.Println("cancel")
		return errors.New("cancel")
	case r := <-respC:
		fmt.Println(r)
		return nil
	}
}

func Func3() {
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func3(ctx, wg)
	time.Sleep(time.Second * 2)
	// 触发取消
	cancel()
	// 等待goroutine退出
	wg.Wait()
}

func func4(ctx context.Context) {
	hctx, hcancel := context.WithTimeout(ctx, time.Second*4)
	defer hcancel()

	resp := make(chan struct{}, 1)
	// 处理逻辑
	go func() {
		// 处理耗时
		time.Sleep(time.Second * 5)
		resp <- struct{}{}
	}()

	// 超时机制
	select {
	case <-ctx.Done():
	   fmt.Println("ctx timeout")
	   fmt.Println(ctx.Err())
	case <-hctx.Done():
		fmt.Println("hctx timeout")
		fmt.Println(hctx.Err())
	case v := <-resp:
		fmt.Println("test2 function handle done")
		fmt.Printf("result: %v\n", v)
	}
	fmt.Println("test2 finish")
	return
}

func Func4() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	func4(ctx)
}






//new one
func Fa1(ctx context.Context){
	ctx = context.WithValue(ctx, 12, "13")
	Fa2(ctx)
}

func Fa2(ctx context.Context){
	fmt.Println(ctx.Value(12).(string))
}

// 耗时场景，通过select模拟
func Fa3() error{
	respC := make(chan int)

	go func(){
		time.Sleep(time.Second*2)
		respC <- 10
		close(respC)
	}()

	select {
	case r:=<-respC:
		fmt.Printf("resp: %d\n", r)
		return nil
	case <-time.After(time.Second*2):
		fmt.Println("catch timeout")
	return errors.New("timeout")
	}
}

// 主动取消
func Fa4(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	respC := make(chan int)

	go func() {
		time.Sleep(time.Second*2)
		respC <- 11
	}()

	select {
	case <- ctx.Done():
		fmt.Println("cancel")
		return errors.New("cancel")
	case r := <-respC:
		fmt.Println(r)
		return nil
	}
}

// 超时取消
func Fa5(ctx context.Context){
	hctx, hcancel := context.WithTimeout(ctx, time.Second*4)
	defer hcancel()

	resp := make(chan struct{}, 1)

	go func(){
		time.Sleep(time.Second*30)
		resp <- struct{}{}
	}()

	select {
	case <- ctx.Done():
		fmt.Println("ctx timeout")
		fmt.Println(ctx.Err())
		return
	case <- hctx.Done():
		fmt.Println("htcx timeout")
		fmt.Println(hctx.Err())
	case v := <- resp:
		fmt.Println("test2 function handle done")
		fmt.Printf("result: %v\n", v)
	}
	fmt.Println("test2 finish")
	return
}





func main2(){
	//创建trace文件
    f, err := os.Create("trace.out")
    if err != nil {
        panic(err)
    }

    defer f.Close()

    //启动trace goroutine
    err = trace.Start(f)
    if err != nil {
        panic(err)
    }
    defer trace.Stop()


	// ctx := context.Background()
	// Fa1(ctx)

	// err1 := Fa3()
	// fmt.Printf("fa3 errors: %v\n", err1)

	// wg := new(sync.WaitGroup)
	// ctx, cancel := context.WithCancel(context.Background())
	// wg.Add(1)
	// go Fa4(ctx, wg)
	// time.Sleep(time.Second*2)
	// cancel()
	// wg.Wait()


	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	Fa5(ctx)
}



type Context interface{
	Deadline()
	Done()
	Err()
	Value()
}

type emptyCtx int

func (*emptyCtx) Deadline()(deadline time.Time, ok bool){
	return
}

func (*emptyCtx) Done() <-chan struct{}{
	return nil
}

func (*emptyCtx) Err() error{
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{}{
	return nil
}


type cancelContext struct {
	emptyCtx
	mu sync.Mutex
	done chan struct{}
	children map[canceler]struct{}
	err error
}

type canceler interface{
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}


