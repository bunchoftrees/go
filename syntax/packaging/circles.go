package main

import "math"


func Truncate(num float64) float64{
	var truncated = math.Round(num * 100) / 100
	return truncated
}

func Circumference(radius float64) float64 {
	var circ = 2 * math.Pi * radius
	circ = Truncate(circ)
	return circ
}

func Area(radius float64) float64 {
	var area = math.Pi * math.Pow(radius, 2)
	area = Truncate(area)
	return area
}
