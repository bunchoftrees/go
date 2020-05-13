package measurements

import "math"

// Function truncate is unexported
func truncate(num float64) float64 {
	var truncated = math.Round(num*100) / 100
	return truncated
}

/* Circumference is a function that can be exported.
Calculates circumference using radius of type float64, returns
the circumference rounded to the nearest hundredth.
*/
func Circumference(radius float64) float64 {
	var circ = 2 * math.Pi * radius
	circ = truncate(circ)
	return circ
}

func AreaCircle(radius float64) float64 {
	var area = math.Pi * math.Pow(radius, 2)
	area = truncate(area)
	return area
}
