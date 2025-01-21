package main

import (
	"fmt"
	// "math"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	A = iota
	B = iota
	C
	D = -iota
	E
	F = "ad"
	G
	H = 100
	I
	J = iota
	K
)

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
	bb := s[0]
	fmt.Println(bb)
	if len(s) <= 1 {
		fmt.Println(s)
	}
	for i := 1; i<len(s);i++{
		cc := s[i]
		fmt.Println(cc)
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


	p1 := 10
	p1s := strconv.Itoa(p1)
	p2, _ := strconv.Atoi(p1s)
	p3, _ := strconv.ParseInt(p1s, 10, 8)
	p4s := strconv.FormatInt(p3, 10)
	fmt.Println(p1,p1s,p2,p3,p4s)

}

type Phone interface{
	call()
}

type iPhone struct{}
type pPhone struct{}

func (i iPhone) call() {
	fmt.Println("iPhone")
}

func (p pPhone) call() {
	fmt.Println("pPhone")
}

func testInterface(){
	var p Phone
	p = iPhone{}
	p.call()
	p = pPhone{}
	p.call()
}

func testForRange(){
	arr := []int{1,2,3,4}
	for a, b := range arr {
		fmt.Println(a, &a, b, &b, &arr[a])
	}
}

func testReflect(){
	type MyStruct struct {
		A int
		B string
	}

	s := MyStruct{A: 20, B: "hello"}
	sv := reflect.ValueOf(s)

	for i := 0; i < sv.NumField(); i++ {
		valueField := sv.Field(i)
		typeField := sv.Type().Field(i)
		fmt.Printf("Field Name: %s, Field Value: %v\n", typeField.Name, valueField.Interface())
	}
}

func testMap(){
	a := make(map[int]int, 0)
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4
	a[5] = 5
	for i, k := range a {
		fmt.Println(i, k)
	}

	b := make(map[chan int]int, 0)
	fmt.Println(b)
}

func testGo() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		select {
		case <- time.After(1000*time.Second):
		}
		wg.Done()
	}()

	go func() {
		select {
		case <- time.After(1000*time.Second):
		}
		wg.Done()
	}()
	wg.Wait()
}

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
	// teststring()
	// testInterface()

	// testForRange()
	// testReflect()
	// testMap()

	// testGo()
	// tt()

	// a := " s s s "
	// b :=[]byte(a)
	// fmt.Println(b)
	// c := string(b)
	// fmt.Println(c)
	// if b[0] == 's' {
	// 	fmt.Println("s")
	// }
	// a = strings.TrimSpace(a)
	// fmt.Println(a, len(a))
	// grid := [][]byte{}
	// isOver(grid, 0, 0)
	// fmt.Println(reflect.TypeOf(a[0]))

	// mp := map[uint8]int{}
	// mp[a[0]]++
	// mp[a[1]]++
	// fmt.Println(mp)
	// delete(mp, a[0])
	// fmt.Println(mp)
	// // strings.Split(a, ",")

	// aaa := 1.1
	// baa := math.Sqrt(aaa)
	// fmt.Println(baa)
	num1 := []int{1,2,3}
	num2 := []int{4,5}
	num3 := copy(num2, num1)
	fmt.Println(num3, num1, num2)
}

func isOver(grid [][]byte, x, y int) bool {
    return x>=0 && x<len(grid) && y>=0 && y<len(grid[0])
}

func tt() {
	a := []int{1, 2, 3}
	fmt.Println(cap(a))
	b := []int{2, 3, 4, 5, 6}
	c := []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a = append(a, b...)
	fmt.Println(cap(a))
	// a = append(a, 3)
	fmt.Println(cap(a))
	a = append(a, c...)
	fmt.Println(cap(a))
	for i:=0;i<1000;i++{
		a =append(a, i)
	}
	fmt.Println(cap(a))
	// a = a[:0]
	// fmt.Println(cap(a))
	d := a[8:20]
	fmt.Println(cap(d))
	e := a[:20]
	fmt.Println(cap(e))
	f := a[10:]
	fmt.Println(cap(f))


	ch := make(chan int, 5)
	ch<- 5
	ch<- 4
	close(ch)
	s := <-ch
	fmt.Println(s)
	s = <-ch
	fmt.Println(s)
	s = <-ch
	fmt.Println(s)
}

