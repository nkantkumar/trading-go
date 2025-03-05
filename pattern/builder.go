package main

import "fmt"

type House struct {
	Floors  int
	Color   string
	HasPool bool
}

type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) SetFloors(floors int) *HouseBuilder {
	b.house.Floors = floors
	return b
}

func (b *HouseBuilder) SetColor(color string) *HouseBuilder {
	b.house.Color = color
	return b
}

func (b *HouseBuilder) SetPool(pool bool) *HouseBuilder {
	b.house.HasPool = pool
	return b
}

func (b *HouseBuilder) Build() House {
	return b.house
}

func main() {
	builder := &HouseBuilder{}
	house := builder.SetFloors(2).SetColor("Blue").SetPool(true).Build()
	fmt.Printf("%+v\n", house)
}
