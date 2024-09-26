// Garbage Collected

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func new(name string, age int) *Person {
	person := Person{
		name: name,
		age:  age,
	}
	return &person
}

func main() {
	jio := new("Jio", 23)
	cla := new("Cla", 23)

	persons := make([]Person, 2)
	persons[0] = *jio
	persons[1] = *cla
	fmt.Printf("%+v %d\n", persons, len(persons))

	sayHello(*jio)
	sayHello(*cla)

	fmt.Printf("%s, %d\n", jio.name, jio.age)
	fmt.Printf("%s, %d\n", cla.name, cla.age)
}

func sayHello(person Person) {
	fmt.Printf("Hello %s!\n", person.name)
}
