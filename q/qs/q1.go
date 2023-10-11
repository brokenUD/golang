package qs

import "fmt"

func Q1() {
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 1)
	chan3 := make(chan int, 1)

	go func() {
		for i := 0;i < 26 ;i ++ {
			chan2<-i
			fmt.Printf("%v\n", i+1)
			<-chan1
		}
	}()

	go func() {
		for i := 0;i < 26 ;i ++ {
			<-chan2
			fmt.Printf("%c\n", 'a'+i)
			chan1<-i
		}
		chan3<-1
	}()
	<-chan3
	close(chan1)
	close(chan2)
	close(chan3)

	name := point()
    fmt.Println(*name)

	chan4 := make(chan int, 4)
	chan4 <- 1
	chan4 <- 2
	close(chan4)
	// for k:= range chan4 {
	// 	fmt.Println(k)
	// }
	k1, v1 := <-chan4
	fmt.Println(k1, v1)
	k2, v2 := <-chan4
	fmt.Println(k2, v2)
	k3, v3 := <-chan4
	fmt.Println(k3, v3)
}


func point() *string {
    name := "指针"
    return &name
}
