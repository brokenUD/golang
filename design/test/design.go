package test

import (
	"fmt"
	// "math/rand"

	// "golang.org/x/exp/rand"
	// "time"
	// "strconv"
)

type Fruit interface{
	Eat()
}

type Lemon struct{
	name string
}

func NewLemon(name string) Fruit{
	return &Lemon{
		name: name,
	}
}

func (l *Lemon)Eat(){
	fmt.Printf("lemon %s", l.name)
}






