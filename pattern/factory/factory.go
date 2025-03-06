package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

func AnimalFactory(animalType string) Animal {
	if animalType == "dog" {
		return Dog{}
	} else if animalType == "cat" {
		return Cat{}
	}
	return nil
}

func main() {
	animal := AnimalFactory("dog")
	fmt.Println(animal.Speak()) // Woof!
}
