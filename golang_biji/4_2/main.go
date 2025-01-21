package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func main(){
	file, err := os.Create("./cpu.pprof")
	if err != nil {
		fmt.Println(err)
	}

	pprof.StartCPUProfile(file)

	defer file.Close()
	defer pprof.StopCPUProfile()

	for i:=0;i<10;i++{
		go actionSort()
	}
	time.Sleep(time.Second*10)
}

func actionSort() {
	rand.Seed(time.Now().Unix())
	number := make([]int, 10000, 30000)
	for k, _ := range number {
		number[k] = rand.Intn(1000)
	}

	L := len(number)
	BubbleSort(number, L)
}

func BubbleSort(data []int, L int){
	t := time.Now()
	for i:=0;i<L;i++{
		for j:=i+1;j<L;j++{
			if data[i]>data[j]{
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	fmt.Println("BubbleSort - Time:", time.Since(t))
}

