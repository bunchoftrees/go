package main

import (
	"fmt"
	"github.com/bunchoftrees/go/syntax/packaging/greet"
	"github.com/bunchoftrees/go/syntax/packaging/measurements"
)

func main() {
	greet.SayHello()
	var myRadius float64 = 5

	fmt.Printf("The area of a circle with the radius %v is %v\n", myRadius, measurements.AreaCircle(myRadius))
	fmt.Printf("The circumference of a circle with the radius %v is %v\n", myRadius,
		measurements.Circumference(myRadius))

	// TODO: refactor using a struct to define a rectangle (https://www.golang-book.com/books/intro/9)
	var recHeight float64 = 3
	var recWidth float64 = 4
	recArea := measurements.AreaRectangle(recHeight, recWidth)
	recPerimeter := measurements.Perimeter(recHeight, recWidth)
	recDiagonal := measurements.Diagonal(recHeight, recWidth)
	fmt.Printf("The area of a rectangle of height %v and width %v is %v\n", recHeight, recWidth, recArea)
	fmt.Printf("The perimeter of a rectangle of height %v and width %v is %v\n", recHeight, recWidth, recPerimeter)
	fmt.Printf("The diagonal of a rectangle of height %v and width %v is %v\n", recHeight, recWidth, recDiagonal)
}
