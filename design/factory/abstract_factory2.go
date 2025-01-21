package factory

import (
	"fmt"
)

type Strawberry3 interface {
	SweetAttack()
}

type Lemon3 interface {
	AcidAttack()
}

type FruitFactory3 interface {
	CreateStrawberry() Strawberry3
	CreateLemon() Lemon3
}

type GoodfarmerStrawberry struct {
	brand string
	Strawberry3
}

func (g *GoodfarmerStrawberry) SweetAttack() {
	fmt.Printf("sweet %s", g.brand)
}

type GoodfarmerLemon struct {
	brand string
	Lemon3
}

func (g *GoodfarmerLemon) AcidAttack() {
	fmt.Printf("acid %s", g.brand)
}

type DolefarmerStrawberry struct {
	brand string
	Strawberry3
}

func (d *DolefarmerStrawberry) SweetAttack() {
	fmt.Printf("sweet %s", d.brand)
}

type DolefarmerLemon struct {
	brand string
	Lemon3
}

func (d *DolefarmerLemon) AcidAttack() {
	fmt.Printf("acid %s", d.brand)
}


type GoodfarmerFactory struct{}

func (g *GoodfarmerFactory)myAspect(){
	fmt.Println("good farmer")
}

func (g *GoodfarmerFactory)CreateStrawberry() Strawberry3{
	g.myAspect()
	return &GoodfarmerStrawberry{
		brand: "good farmer",
	}
}

func (g *GoodfarmerFactory)CreateLemon() Lemon3{
	g.myAspect()
	return &GoodfarmerLemon{
		brand: "good farmer",
	}
}

type DoleFactory struct{}

func (d *DoleFactory)myAspect(){
	fmt.Println("Dole farmer")
}

func (d *DoleFactory)CreateStrawberry() Strawberry3{
	d.myAspect()
	return &DolefarmerStrawberry{
		brand: "Dole farmer",
	}
}

func (d *DoleFactory)CreateLemon() Lemon3{
	d.myAspect()
	return &DolefarmerLemon{
		brand: "Dole farmer",
	}
}
