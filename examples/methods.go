package main

import "fmt"

type rect struct {
	width, height int
}

// pointer receiver when? --> you need to modify the struct or it is so big and you don't want to copy every call (performance)
func (r *rect) area() int {
	return r.width * r.height
}

// value receiver when? --> read only with small struct or you want to be sure that the struct will not edit
func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter: ", rp.perimeter())

	// Go automatically handles conversion between values and pointers for method calls
}