package main

import (
	"fmt"
	"github.com/bunchoftrees/go/syntax/types/caster"
	"reflect"
)

func main() {
	// Integer
	fmt.Println("Type: ", reflect.TypeOf(1))
	// Float (64)
	fmt.Println("Type: ", reflect.TypeOf(1.5))
	// String
	fmt.Println("Type: ", reflect.TypeOf("Hello"))
	// Char (sort of)
	fmt.Println("Type: ", reflect.TypeOf('c'))
	// Boolean
	fmt.Println("Type: ", reflect.TypeOf(true))

	// What type is this? ;)
	whatType := caster.Cast(5)
	fmt.Println("Type: ", reflect.TypeOf(whatType))
}
