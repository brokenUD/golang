package main

import (
	"fmt"
	"sync"
	// "time"
	// "sync"
)

func main() {
	ch1 := make(chan int)
	defer close(ch1)
	ch2 := make(chan int)
	defer close(ch2)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func(){
		defer wg.Done()
		i := 0
		for {
			fmt.Println(string('A'+i))
			ch1 <- i
			<- ch2
			i++
			if i >= 26 {
				return
			}
		}
		// wg.Done()
	}()

	go func(){
		defer wg.Done()
		i := 0
		for {
			<- ch1
			fmt.Println(i)
			ch2 <- i
			i++
			if i >= 26 {
				return
			}
		}
		// wg.Done()
	}()

	wg.Wait()
	// time.Sleep(time.Second)
	// return
}
