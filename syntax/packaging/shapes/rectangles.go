package shapes

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width, Length float64
}

func NewRectangle(width float64, length float64) *Rectangle {
	return &Rectangle{Width: width, Length: length}
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (r *Rectangle) Perimeter() float64 {
	return (r.Width * 2) + (r.Length * 2)
}

func (r *Rectangle) Diagonal() float64 {
	return math.Sqrt(math.Pow(r.Width, 2) + math.Pow(r.Length, 2))
}

func (r *Rectangle) IsSquare()  {
	if r.Width == r.Length {
		fmt.Println("This rectangle is a square!")
	} else {
		fmt.Println("This rectangle is not a square.")
	}
}
