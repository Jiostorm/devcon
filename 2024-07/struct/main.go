package main

import "fmt"

type Cat struct {
	Name  string
	Owner string
}

func (cat *Cat) sound() string {
	return fmt.Sprintf("%s says meooooow!", cat.Name)
}

func (cat *Cat) climb() string {
	return fmt.Sprintf("%s climbed the wall!", cat.Name)
}

type Dog struct {
	Name  string
	Owner string
}

func (dog *Dog) sound() string {
	return fmt.Sprintf("%s says hooooowl!", dog.Name)
}

type Lizard struct {
	Name  string
	Owner string
}

func (lizard *Lizard) climb() string {
	return fmt.Sprintf("%s climbed the wall!", lizard.Name)
}

type Pet interface {
	sound() string
}

type Climber interface {
	climb() string
}

func main() {
	pets := []Pet{&Cat{Name: "Kiel", Owner: "Jio-"}, &Dog{Name: "Hazel", Owner: "Issa"}}
	climbers := []Climber{
		&Cat{Name: "Kiel", Owner: "Jio-"},
		&Lizard{Name: "Lizardo", Owner: "Cardo"},
	}

	for _, pet := range pets {
		fmt.Println(pet.sound())
	}

	for _, climber := range climbers {
		fmt.Println(climber.climb())
	}
}
