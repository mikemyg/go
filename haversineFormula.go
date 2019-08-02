package main

import (
	"math"
)

func harvesineFormula(lat1 float64,
	lon1 float64,
	lat2 float64,
	lon2 float64) float64 {

	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0

	radianslat1 := (lat1) * math.Pi / 180.0
	radianslat2 := (lat2) * math.Pi / 180.0

	a := math.Pow(math.Sin(dLat/2), 2) +
		math.Pow(math.Sin(dLon/2), 2)*
			math.Cos(radianslat1)*math.Cos(radianslat2)

	var rad float64
	rad = 6371
	c := 2 * math.Asin(math.Sqrt(a))
	c = c * rad

	//return distance in km
	return c
}
