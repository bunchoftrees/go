package main

import (
	"fmt"
	"github.com/bunchoftrees/go/syntax/types/caster"
	"net"
	"reflect"
	"time"
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
	// net.IPv4
	fmt.Println("Type: ", reflect.TypeOf(net.IPv4(127, 0, 0, 1)))
	// time.Time
	fmt.Println("Type: ", reflect.TypeOf(time.Now()))

	// What type is this? ;)
	whatType := caster.Cast(5)
	fmt.Println("Type: ", reflect.TypeOf(whatType))
}
