package adapt

import (
	"fmt"
)

type inter1 interface {
	Eat()
}

type A struct {}

func (a *A)Eat(){
	fmt.Println("A eat")
}

type inter2 interface {
	Eatten()
}

type B struct {}

func (b *B)Eatten(){
	fmt.Println("B eatten")
}

type Adapter struct {
	b *B
}

func (a *Adapter)Eat(){
	a.b.Eatten()
}
