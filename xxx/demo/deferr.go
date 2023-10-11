package demotest

import "fmt"

func DemotestFun1() {
    var whatever [5]struct{}
    for i := range whatever {
        defer fmt.Println(i)
    }
}


func DemotestFun2() {
	i := 1
	fmt.Println("i =", i)
	defer fmt.Print(i)
}



func DemotestFun3(a int) {//无返回值函数
	defer fmt.Println("1、a =", a) //方法
	defer func(v int) { fmt.Println("2、a =", v)} (a) //有参函数
	defer func() { fmt.Println("3、a =", a)} () //无参函数
	a++
}

func DemotestFun4() {//无返回值函数
	DemotestFun3(4)
}


func DemotestFun5(i int) int{
	return i
	defer fmt.Print("i =", i)
	return i+1
}

func DemotestFun6() {
	DemotestFun5(1)
}






func DemotestFun7() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

func DemotestFun8() {
	fmt.Println("return:", DemotestFun7())
}






func DemotestFun9() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func DemotestFun10() {
	fmt.Println("return:", DemotestFun9())
}










func DemotestFun11() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

func DemotestFun12() {
	fmt.Println("return:", *(DemotestFun11()))
}




func DemotestFun13() (r int) {
    defer func(r int) {
          r = r + 5
    }(r)
    return 1
}

func DemotestFun14() {
	fmt.Println("return:", DemotestFun13())
}





type Test struct {
	name string
}
func (t *Test) pp() {
	fmt.Println(t.name)
}
func DemotestFun15() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.pp()
	}
}






func pp2(t Test) {
	fmt.Println(t.name)
}
func DemotestFun16() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer pp2(t)
	}
}










func DemotestFun17() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		tt := t
		println(&tt)
		defer tt.pp()
	}
}







func DemotestFun18() {
    panic("panic")
    defer fmt.Println("defer after panic")
}



func DemotestFun19() {
    defer fmt.Println("defer before panic")
    panic("panic")
}





