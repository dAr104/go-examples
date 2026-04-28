package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Go string literals are UTF-8 encoded text
	const s = "สวัสดี"

	// A go string is a read-only slices of bytes. len(s) = # bytes
	fmt.Println("Len:", len(s))

	// Indexing a string produces the raw bytes
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()

	// in other languages the strings are made by characters. In go the concept of character is a "Rune", a int32 that rappresent a Unicode code point
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't'{
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}