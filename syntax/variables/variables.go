package main

import "fmt"

func main() {
	var a, b int
	a, b = 1, 2
	var c, d = 3, 4
	e, f := 5, 6
	fmt.Println(a, b, c, d, e, f)
	a, b = 3, 4
	fmt.Println(a, b)
	// Oops, better comment these out - we can't reassign types to variables
	//c, d = 3.1, 4.1
	//fmt.Println(c, d)

	// Naming conventions
	word := "Hello"
	multipleWords := "Hello World"
	// These will work, but it's best practice to not name variables like these
	// multiplewords
	// multiple_words

	fmt.Println(word, multipleWords)
}
