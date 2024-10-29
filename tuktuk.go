package main

import "math"

func calculateFare(km float64, minutes float64) float64 {
	return 4.0*km + minutes
}

func roundUpDistance(km float64) float64 {
	return math.Ceil(km*2) / 2
}
