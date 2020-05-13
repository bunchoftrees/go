package shapes

import "math"

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

// Function truncate is unexported
func truncate(num float64) float64 {
	var truncated = math.Round(num*100) / 100
	return truncated
}

func (c *Circle) Circumference() float64 {
	var circ = 2 * math.Pi * c.Radius
	circ = truncate(circ)
	return circ
}

func (c *Circle) Area() float64 {
	var area = math.Pi * math.Pow(c.Radius, 2)
	area = truncate(area)
	return area
}
