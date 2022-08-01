package contextt

import (
	"context"
	"errors"
	"fmt"
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
		time.Sleep(time.Second * 3)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	func4(ctx)
}
