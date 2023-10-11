package demotest

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	n int32
	lock = sync.Mutex{}
)

func foo(){
	for i:=0;i<100000;i++{
		lock.Lock()
		n++
		lock.Unlock()
	}
	fmt.Println("n=",n)
}

func BfTest(){
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		defer wg.Done()
		foo()
	}()
	go func(){
		defer wg.Done()
		foo()
	}()
	wg.Wait()
	fmt.Println("main n=", n)
}


func foo2(){
	for i:=0;i<100000;i++{
		atomic.AddInt32(&n, 1)  //原子操作
	}
	fmt.Println("n=",n)
}

func BfTest2(){
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		defer wg.Done()
		foo2()
	}()
	go func(){
		defer wg.Done()
		foo2()
	}()
	wg.Wait()
	fmt.Println("main n=", n)
}
