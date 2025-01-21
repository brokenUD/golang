package factory

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Fruit interface {
	Eat()
}

type Orange struct {
	name string
}

func NewOrange(name string) Fruit {
	return &Orange{
		name: name,
	}
}

func (o *Orange) Eat() {
	fmt.Printf("i am %s", o.name)
}

type Strawberry struct {
	name string
}

func NewStrawberry(name string) Fruit {
	return &Strawberry{
		name: name,
	}
}

func (s *Strawberry) Eat() {
	fmt.Printf("i am %s", s.name)
}

type Cherry struct {
	name string
}

func NewCherry(name string) Fruit {
	return &Cherry{
		name: name,
	}
}

func (c *Cherry) Eat() {
	fmt.Printf("i am %s", c.name)
}

// type FruitFactory struct{}

// func NewFruitFactory()*FruitFactory{
// 	return &FruitFactory{}
// }

// func (f *FruitFactory)CreateFruit(typ string)(Fruit, error){
// 	src := rand.NewSource(time.Now().UnixNano())
// 	rander := rand.New(src)
// 	name := strconv.Itoa(rander.Int())

// 	switch typ{
// 	case "orange":
// 		return NewOrange(name), nil
// 	case "strawberry":
// 		return NewStrawberry(name), nil
// 	case "cherry":
// 		return NewCherry(name), nil
// 	default:
// 		return nil, fmt.Errorf("fruit typ %s not supported", typ)
// 	}
// }

type fruitCreator func(name string) Fruit

type FruitFactory struct {
	creators map[string]fruitCreator
}

func NewFruitFactory() *FruitFactory {
	return &FruitFactory{
		creators: map[string]fruitCreator{
			"orange":     NewOrange,
			"strawberry": NewStrawberry,
			"cherry":     NewCherry,
		},
	}
}

func (f *FruitFactory)CreateFruit(typ string)(Fruit, error){
	fruitCreator, ok := f.creators[typ]
	if !ok {
		return nil, fmt.Errorf("type %s not", typ)
	}

	src := rand.NewSource(time.Now().Unix())
	rander := rand.New(src)
	name := strconv.Itoa(rander.Int())
	return fruitCreator(name), nil
}
