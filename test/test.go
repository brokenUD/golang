package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	// var bool1 bool
	// var string1 string
	// flag1 := flag.Bool("bool1", false, "布尔值")
	// flag.Parse()
	// fmt.Printf("%v, %v", flag.Args(), flag.NArg())
	// flag.StringVar(&string1, "string1", "hahaha", "字符串")
	// fmt.Println(*flag1)
	// fmt.Println(string1)

	// test1()
	// test2()
	// testchan()
	// sliceSafety()
	// syncMapTest()

	// testtest()

	// testsort()
	teststring()
}

func test1() {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		a[k] = 100 + v
	}
	fmt.Print(a)
}

func test2() {
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		a[k] = 100 + v
	}
	fmt.Print(a)
}

func testchan() {
	c := make(chan int)
	defer close(c)
	go func() { c <- 3 + 4 }()
	i := <-c
	fmt.Println(i)

	k, v := <-c
	fmt.Println(k)
	fmt.Println(v)
}

var m sync.Mutex
var w sync.WaitGroup

func sliceSafety() {
	var s []int
	var sum int
	fmt.Printf("----------: len(s): %d, cap(s): %d, s: %v \n", len(s), cap(s), s)
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			m.Lock()
			sum++
			s = append(s, i)
			fmt.Printf("==========i: %d: len(s): %d, cap(s): %d, s: %v \n", i, len(s), cap(s), s)
			m.Unlock()
		}(i)
	}
	w.Wait()
	fmt.Println(sum)
	fmt.Println(s, len(s))
}

var sm sync.Map

func syncMapTest() {
	//写入
	data := []string{"煎鱼", "咸鱼", "烤鱼", "蒸鱼"}
	for i := 0; i < 4; i++ {
		go func(i int) {
			sm.Store(i, data[i])
		}(i)
	}
	time.Sleep(time.Second)

	//读取
	v, ok := sm.Load(4)
	fmt.Printf("Load: %v, %v\n", v, ok)

	//删除
	sm.Delete(4)

	//读或写
	v, ok = sm.LoadOrStore(1, "吸鱼")
	fmt.Printf("LoadOrStore: %v, %v\n", v, ok)

	//遍历
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("Range: %v, %v\n", key, value)
		return true
	})
}

func testtest() {
	fmt.Printf("%b", 39082)
	rune1 := []rune{20320, 22909, 12312}
	byte1 := []byte{220, 209, 112}
	fmt.Println(rune1)
	fmt.Println(string(rune1))
	fmt.Println(string(byte1))
	// s1 := "你好"
	// for k, v := range s1 {
	// 	fmt.Println(k, v)
	// }
	arr1 := [5]int{2,3}
	slice1 := []int{2,3}
	fmt.Printf("%p, %p\n", &arr1, &arr1[0])
	fmt.Printf("%p, %p\n", &slice1, &slice1[0])

	var chan1 chan int
	chan1 = make(chan int, 10)
	chan1 <- 1
	chan1 <- 2
	for i:=0;i<8;i++{
		chan1<-3
	}
	fmt.Println(len(chan1))
	close(chan1)
	for v:=range chan1{
		fmt.Println(v,len(chan1))
	}


}

func testsort(){
	type atype struct{
		num int
	}
	var a []atype
	a = append(a, atype{num:2}, atype{num:4} , atype{num:1})
	sort.Slice(a , func(i,j int)bool{
		return a[i].num>a[j].num
	})
	fmt.Println(a)
}

func teststring(){
	s := "qwertresfs"
	max := ""
	if len(s) <= 1 {
		fmt.Println(s)
	}
	for i := 1; i<len(s);i++{
		if s[i] == s[i-1] {
			min := ""
			p := i-1
			q := i
			for (p>0 && q<len(s)-1 && s[p-1]==s[q+1]){
				p -= 1
				q += 1
			}
			min = s[p:q+1]
			if len(min) > len(max){
				max = min
			}
		} else if i>2 && s[i] == s[i-2] {
			min := ""
			p := i-2
			q := i
			for (p>0 && q<len(s)-1 && s[p-1]==s[q+1]){
				p--
				q++
			}
			min = s[p:q+1]
			if len(min) > len(max){
				max = min
			}
		}
	}
	fmt.Println(max)

	// a := "ssss"
	// b := len(a)
	s0 := "1234567890"
	// var s1 strings.Builder
	// s1.Grow(11)
	// s1.WriteString(s0[:10])
	// s2 := s1.String()
	// fmt.Println(s1, s0, s2, &s1, &s2, &s0)

	s3 := strings.Repeat(s0, 2)
	fmt.Println(s0, s3, &s0, &s3)

	// a := [3]int{}


	q1:= "asdfghj"
	fmt.Println(strings.HasPrefix(q1, "s"))
	fmt.Println(strings.HasSuffix(q1, "s"))


	num1, num2 := -5, 2
	fmt.Println(num1/num2)


	qq1 := []byte{}
	qq1 = append(qq1, 'a')
	if qq1 == q1[2:3] {
		fmt.Println(true)
	}
	for i:=0;i<len(q1);i++{
		t := q1[i]
		fmt.Println(t)
	}
}


