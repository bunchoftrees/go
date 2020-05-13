package measurements

import (
	"math"
)

func AreaRectangle(x float64, y float64) float64 {
	return x * y
}

func Perimeter(x float64, y float64) float64 {
	return x + x + y + y
}

func Diagonal(x float64, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
