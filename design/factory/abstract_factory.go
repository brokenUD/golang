package factory

import (
	// "fmt"
)

type FruitFactory2 interface {
	CreateFruit() Fruit
}

type OrangeFactory struct{}

func NewOrangeFactory() FruitFactory2{
	return &OrangeFactory{}
}

func (o *OrangeFactory)CreateFruit()Fruit{
	return NewOrange("")
}

type StrawberryFactory struct{}

func NewStrawberryFactory() FruitFactory2{
	return &StrawberryFactory{}
}

func (o *StrawberryFactory)CreateFruit()Fruit{
	return NewStrawberry("")
}

type CherryFactory struct{}

func NewCherryFactory() FruitFactory2{
	return &CherryFactory{}
}

func (o *CherryFactory)CreateFruit()Fruit{
	return NewCherry("")
}
