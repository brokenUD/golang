package factory

import "testing"

func Test_factory(t *testing.T){
	FruitFactory := NewFruitFactory()

	orange, _ := FruitFactory.CreateFruit("orange")
	orange.Eat()

	cherry, _ := FruitFactory.CreateFruit("cherry")
	cherry.Eat()
}
