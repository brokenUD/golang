package demotest

import (
	"fmt"
	"time"
	"sync"
)

func Thirt(a int){
	for i:=0;i<a;i++{
		fmt.Println("third")
	}
}

func f1(){
	fmt.Println("f1")
	go f2()
}

func f2(){
	time.Sleep(1*time.Second)
	fmt.Println("f2")
}

func GooTest(){
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func(){
		defer wg.Done()
		fmt.Println("first")
	}()
	go func(){
		defer wg.Done()
		fmt.Println("second")
	}()
	go Thirt(3)
	wg.Done()
	wg.Wait()


	go f1()
	fmt.Println("f1 finish")
	time.Sleep(2*time.Second)
}

func f3(){
	defer func(){  // 不能写成defer recover()
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	a, b := 3, 0
	fmt.Println(a, b)
	_ = a/b
	fmt.Println("f3 finish")
}

func RecoverTest(){
	go f3()
	fmt.Println("f3")
	time.Sleep(1*time.Second)
	fmt.Println("over")
}
