package demotest

import (
	"fmt"
	"sync"
	// "time"
)

func DeathLockTest(){
	ch := make(chan int, 0)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(){
		defer wg.Done()
		ch <- 1
		fmt.Println("over")
	}()

	wg.Wait()
}
