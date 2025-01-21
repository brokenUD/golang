package main

import (
	"fmt"
)

type ai interface {
	Eat()
}

type A struct {}

func (a *A)Eat(){
	fmt.Println("a eat")
}

type Decor ai

func NewAA(aa ai)Decor {
	return aa
}

type lAA struct {
	Decor
}

func (l *lAA)Eat(){
	fmt.Println("laa")
	l.Decor.Eat()
}

func main() {
	b := &A{}
	a := NewAA(b)
	c := &lAA{
		Decor: a,
	}
	// a.AA =
	c.Eat()
}
