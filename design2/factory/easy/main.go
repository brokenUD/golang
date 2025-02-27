package main

import "fmt"

type vehicle interface {
	Drive() string
}

type car struct{}

func (c *car) Drive() string {
	return "drive car"
}

type bike struct{}

func (b *bike) Drive() string {
	return "ride bike"
}

func GetVehicle(typ string) vehicle {
	switch typ {
	case "car":
		return &car{}
	case "bike":
		return &bike{}
	default:
		return nil
	}
}

func main() {
	a := GetVehicle("bike")
	if a != nil {
		fmt.Println(a.Drive())
	}
	b := GetVehicle("s")
	if b != nil {
		fmt.Println(b.Drive())
	}
}
