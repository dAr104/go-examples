package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// multiple return values e.g. return born result and error values from a function
func vals() (int, int) {
	return 3, 7
}

// Variadic functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// Closures
func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3) 
	fmt.Println("1+2+3 =", res)

	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// test closures
	nextInt := initSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := initSeq()
	fmt.Println(newInts())
}