package main

import (
	"fmt"
	"math/rand"
	"net/http"
	// "os"
	_ "net/http/pprof"
	"time"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		for i:=0;i<10;i++{
			actionSort(w)
		}
	})
	http.ListenAndServe(":9000", nil)
}

func actionSort(w http.ResponseWriter) {
	rand.Seed(time.Now().Unix())
	number := make([]int, 30000, 3000000)
	for k, _ := range number {
		number[k] = rand.Intn(1000)
	}

	L := len(number)
	BubbleSort(w, number, L)
}

func BubbleSort(w http.ResponseWriter, data []int, L int){
	t := time.Now()
	for i:=0;i<L;i++{
		for j:=i+1;j<L;j++{
			if data[i]>data[j]{
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	w.Write([]byte(fmt.Sprintln("BubbleSort - Time:", time.Since(t))))
}

