package main

import "fmt"

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func CreateNewAnimal(a map[string]Animal) {
	var animalName string
	var animalType string

	fmt.Println("Enter animal name: ")
	fmt.Scanln(&animalName)

	fmt.Println("Enter animal type: 'cow', 'bird', 'snake' ")
	fmt.Scanln(&animalType)

	switch animalType {
	case "cow":
		a[animalName] = Cow{}
	case "bird":
		a[animalName] = Bird{}
	case "snake":
		a[animalName] = Snake{}
	}

	fmt.Println("Created it!")
}

func QueryAnimal(a map[string]Animal) {
	var animalName string
	var animalAction string

	fmt.Println("Enter animal name: ")
	fmt.Scanln(&animalName)

	animal, ok := a[animalName]

	if !ok {
		fmt.Println("Can't find name")
		return
	}

	fmt.Println("Enter animal action: 'eat', 'move', 'speak'")
	fmt.Scanln(&animalAction)

	switch animalAction {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("I don't know")
	}
}

func main() {
	var command string
	animals := make(map[string]Animal)

	for {
		fmt.Println("Enter command 'newanimal' or 'query':")
		fmt.Scanln(&command)

		switch command {
		case "newanimal":
			CreateNewAnimal(animals)
		case "query":
			QueryAnimal(animals)
		}
	}
}
