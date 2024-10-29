package main

import "testing"

func TestCalculateFareShouldReturnTotalOf1KmWithoutWaitingTime(t *testing.T) {
	distance := 1.0
	waitingTime := 0.0

	got := calculateFare(distance, waitingTime)

	want := 4.0
	if got != want {
		t.Errorf("expected %f, got %f", want, got)
	}
}
