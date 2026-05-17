package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Pinco"})
	fmt.Println(&person{name: "Fringo", age: 22})
	fmt.Printf("%p\n", &person{name: "Fringo", age: 22})
	fmt.Println(newPerson("Jon")) // it's idiomatic to encapsulate new struct creation in constructor functions

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // dots also on pointers, the pointers are automatically dereferenced

	sp.age = 51
	fmt.Println(sp.age)

	// anonymus struct type
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}