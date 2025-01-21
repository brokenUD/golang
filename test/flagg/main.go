package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

type A struct{
	a int
	b string
}

func main() {
	flag.String("n", "xxx", "name")
	// var a A
	// flag.Var(&A, )
	fmt.Println(os.Getwd())
	// os.Chdir("test/flagg")
	// fmt.Println(os.Getwd())

	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
