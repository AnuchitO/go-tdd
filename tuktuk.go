package main

import "math"

func calculateFare(km float64, minutes float64) float64 {
	return 4.0*km + minutes
}

func roundUpDistance(km float64) float64 {
	return math.Ceil(km*2) / 2
}

func roundUpWaitingTime(seconds int) float64 {
	return math.Ceil(float64(seconds) / 60)
}

func Fare(km float64, seconds int) float64 {
	return calculateFare(roundUpDistance(km), roundUpWaitingTime(seconds))
}

func Minimum(fare float64) float64 {
	return math.Max(35, fare)
}

func Ride(distance float64, waitingTime int) float64 {
	return Minimum(Fare(distance, waitingTime))
}

type ride struct {
	distance    float64
	waitingTime int
}

func TotalFare(rides []ride) float64 {
	var total float64
	for _, ride := range rides {
		total += Ride(ride.distance, ride.waitingTime)
	}
	return total
}

type invoice struct {
	Rides              []ride
	Total              float64
	AverageFarePerRide float64
}

func InvoiceDetails(rides []ride) invoice {
	fare := TotalFare(rides)
	return invoice{
		Rides:              rides,
		Total:              fare,
		AverageFarePerRide: fare / float64(len(rides)),
	}
}
