package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 5)
	fmt.Println("emp:", "len", len(s), "cap", cap(s))

	a := []uint{1, 2, 3, 4}
	fmt.Println("dcl:", "len", len(a), "cap", cap(a))

	// capacity is doubled when appending beyond capacity
	a = append(a, 5)
	fmt.Println("app:", "len", len(a), "cap", cap(a))

	a = append(a, 8)
	fmt.Println("app:", "len", len(a), "cap", cap(a))

	a = append(a, 9, 10, 11)
	fmt.Println("app:", "len", len(a), "cap", cap(a))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	s[3] = "d"
	s[4] = "e"

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	}

	// Slices can be composed into multi-dimensional data structures. 
	// The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}