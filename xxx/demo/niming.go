package demotest

import "fmt"

func optFunc(opt string) func(int, int) int {
	switch opt {
	case "add":
		return func(a int, b int)int {
			return a+b
		}
	case "del":
		return func(a int, b int)int {
			return a-b
		}
	default:
		return func(a int, b int)int {
			return a+b
		}
	}
}

func NimingF(){
	op := optFunc("add")
	fmt.Println(op(1,2))
}
