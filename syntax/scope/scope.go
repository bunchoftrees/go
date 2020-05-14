package main

import "fmt"

var a = "a" // Package-level variable

func main() {
	var b = "b"
	{
		var c = "c"
		{
			var d = "d"
			fmt.Println(a, b, c, d)
		}
		fmt.Println(a, b, c, d) // ERROR: "d" undefined
	}
	fmt.Println(a, b, c, d) // ERROR: "c", "d" undefined
}
