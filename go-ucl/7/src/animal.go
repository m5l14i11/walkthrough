package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func main() {

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	animals := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	var animalType string
	var animalInfo string

	for {
		fmt.Println("Enter type of animal 'cow', 'bird', 'snake'  >")
		fmt.Scanln(&animalType)

		animal, ok := animals[animalType]

		if !ok {
			fmt.Println("I can't recognize this animal")
			continue
		}

		fmt.Println("Enter information 'eat', 'move', 'speak' >")
		fmt.Scanln(&animalInfo)

		if animalInfo == "eat" {
			fmt.Println(animal.Eat())
		} else if animalInfo == "move" {
			fmt.Println(animal.Move())
		} else if animalInfo == "speak" {
			fmt.Println(animal.Speak())
		} else {
			fmt.Println("I don't know")
		}
	}
}
