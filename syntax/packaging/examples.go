package main

import (
	"fmt"
	"github.com/bunchoftrees/go/syntax/packaging/greet"
	"github.com/bunchoftrees/go/syntax/packaging/shapes"
)

func main() {
	greet.SayHello()

	// My Circle
	myCircle := shapes.NewCircle(5)
	fmt.Printf("My circle has a radius of %v units.\n", myCircle.Radius)
	fmt.Printf("The area this circle is %v square units.\n", myCircle.Area())
	fmt.Printf("The circumference this circle is %v units.\n", myCircle.Circumference())
	// TODO: Implement getter for radius

	// My Rectangle
	myRectangle := shapes.NewRectangle(3, 4)
	fmt.Printf("My rectangle has a width of %v units and a length of %v units.\n", myRectangle.Width, myRectangle.Length)
	myRectangle.IsSquare()
	fmt.Printf("The area of this rectangle is %v square units.\n", myRectangle.Area())
	fmt.Printf("The perimeter of this rectangle is %v units.\n", myRectangle.Perimeter())
	fmt.Printf("The diagonal of this rectangle has a length of %v units.\n", myRectangle.Diagonal())
	// TODO: Implement getter for width and length
}
