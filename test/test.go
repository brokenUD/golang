package main

import (
	"flag"
	"fmt"
)

func main() {
	// var bool1 bool
	var string1 string
	flag1 := flag.Bool("bool1", false, "布尔值")
	flag.Parse()
	fmt.Printf("%v, %v", flag.Args(), flag.NArg())
	flag.StringVar(&string1, "string1", "hahaha", "字符串")
	fmt.Println(*flag1)
	fmt.Println(string1)
}