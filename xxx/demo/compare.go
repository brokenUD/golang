package demotest

import (
	"fmt"
	"unsafe"
)


type Orange struct {
	Host string
	Port int
}


func CompareDemo() {

	// 指针
	orange1 := Orange{}
	orange2 := Orange{}

	op11 := &orange1
	op12 := &orange1
	op2 := &orange2

	fmt.Println("op11 == op12 :", op11 == op12) // true
	fmt.Println("op11 == op2 :", op11 == op2) // false

	// 通道
	ch1 := make(chan int, 1)
	ch2 := ch1
	ch3 := make(chan int, 1)

	fmt.Println("ch1 == ch2 :", ch1 == ch2) // true
	fmt.Println("ch1 == ch3 :", ch1 == ch3) // false

	// 结构体
	orange3 := Orange{}
	orange4 := Orange{}
	orange5 := Orange{"host001", 22}

	fmt.Println("orange3 == orange4 :", orange3 == orange4) // true
	fmt.Println("orange3 == orange5 :", orange3 == orange5) // false

	// 数组
	a1 := [1]int{0}
	a2 := [1]int{0}

	fmt.Println("a1 == a2 :", a1 == a2) // true


	var z1 uint
	fmt.Println(unsafe.Sizeof(z1))
	var z2 uint8
	fmt.Println(unsafe.Sizeof(z2))
	var z3 uint16
	fmt.Println(unsafe.Sizeof(z3))
	var z4 uint64
	fmt.Println(unsafe.Sizeof(z4))
	var z5 uintptr
	fmt.Println(unsafe.Sizeof(z5))
	var z6 int
	fmt.Println(unsafe.Sizeof(z6))
	var z7 int8
	fmt.Println(unsafe.Sizeof(z7))
	var z8 string
	fmt.Println(unsafe.Sizeof(z8))
	var z9 bool
	fmt.Println(unsafe.Sizeof(z9))
	var z10 byte
	fmt.Println(unsafe.Sizeof(z10))
	var z11 rune
	fmt.Println(unsafe.Sizeof(z11))
	var z12 Orange
	fmt.Println(unsafe.Sizeof(z12))
	fmt.Printf("%p\n", &z12)
	var z13 Orange
	fmt.Println(unsafe.Sizeof(z13))
	fmt.Printf("%p\n", &z13)
	var z14 struct{}
	fmt.Println(unsafe.Sizeof(z14))
	fmt.Printf("%p\n", &z14)
	var z15 struct{}
	fmt.Println(unsafe.Sizeof(z15))
	fmt.Printf("%p\n", &z15)
	var z16 []string
	fmt.Println(unsafe.Sizeof(z16))
	var z17 []int
	fmt.Println(unsafe.Sizeof(z17))
}

