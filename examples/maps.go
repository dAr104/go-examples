package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["Dario"] = 31
	m["Noemi"] = 28

	fmt.Println("map:", m)

	v1 := m["Dario"]
	v2 := m["Matisse"] // keys doesn't exits --> return zero value

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)

	fmt.Println("Maps len:", len(m))

	delete(m, "Dario") // delete k/v pair
	fmt.Println("map:", m)

	clear(m) // remove all k/v pairs
	fmt.Println("map:", m)

	_, prs := m["Noemi"] // second value return indicates if the key is present
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}