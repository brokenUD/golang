package demotest

import (
	"fmt"
	"strings"
	"strconv"
)

func demoString(a ...string)string{
	return strings.Join(a, "-")
}

func DemoString(){
	s := demoString("ss", "aa", "dd")
	fmt.Println(s)
}


func DemoString1(){
	s := "1223"
	// s := "31241"
	fmt.Println(s)
}

func DemoStrconv(){
	var a int64
	a = 111111
	// b := "sasda"
	c := "222222"

	d, _ := strconv.Atoi(c)
	fmt.Println(d)

	e, _ := strconv.ParseInt(c, 10, 64)
	fmt.Println(e)

	f := strconv.FormatInt(a, 10)
	fmt.Println(f)

}
