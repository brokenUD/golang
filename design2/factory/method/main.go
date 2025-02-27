package main

import "fmt"

type vehicle interface{
	Drive() string
}

type vehicleFactory interface{
	Create() vehicle
}

type CarFactory struct{}
type Car struct{}
func (f *CarFactory)Create()vehicle{return &Car{}}
func (c *Car)Drive()string{return "car drive"}

type BikeFactory struct{}
type Bike struct{}
func (b *BikeFactory)Create()vehicle{return &Bike{}}
func (c *Bike)Drive()string{return "bike ride"}

func main(){
	var f vehicleFactory = &CarFactory{}
	v := f.Create().Drive()
	fmt.Println(v)
}
