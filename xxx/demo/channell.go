package demotest

import (
	"fmt"
	"sync"
	"time"
)

func TimeoutChannelCancel() {
	done := make(chan struct{}, 1)

	go func() {
		done <- struct{}{}
	}()

	select {
	case <- done:
		fmt.Printf("succsse")
		return
	case <- time.After(time.Duration(time.Millisecond*800)):
		fmt.Printf("time out")
		return
	}
}

func ChannelTest(){
	ch := make(chan int, 0)

	go func(){
		ch<-4
		fmt.Println("写入4 ", time.Now().Unix())
	}()

	time.Sleep(2*time.Second)

	go func(){
		a := <-ch
		fmt.Println("取出 ", a, time.Now().Unix())
	}()

	time.Sleep(1*time.Second)
	fmt.Println(time.Now().Unix())
}


func ChannelTest2(){
	ch := make(chan int, 100)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		defer wg.Done()
		for i := 0;i<10;i++ {
			ch<-i
		}
	}()

	go func(){
		defer wg.Done()
		for i := 0;i<10;i++ {
			ch<-i
		}
	}()

	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func(){
		defer wg2.Done()
		sum:=0
		for{
			a, ok := <-ch
			if !ok{    // chan关闭且为空才会ok
				break
			}
			sum+=a
		}
		fmt.Println(sum)
	}()
	wg.Wait()
	close(ch)
	wg2.Wait()
}



func ChannelTest3(){
	ch := make(chan int, 100)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		defer wg.Done()
		for i := 0;i<10;i++ {
			ch<-i
		}
	}()

	go func(){
		defer wg.Done()
		for i := 0;i<10;i++ {
			ch<-i
		}
	}()

	mc := make(chan struct{})
	go func(){
		sum:=0
		for{
			a, ok := <-ch
			if !ok{    // chan关闭且为空才会ok
				break
			}
			sum+=a
		}
		fmt.Println(sum)
		mc<-struct{}{}
	}()
	wg.Wait()
	close(ch)
	<-mc
}
