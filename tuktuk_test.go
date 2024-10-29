package main

import "testing"

func TestCalculateFareShouldReturnTotalOf1KmWithoutWaitingTime(t *testing.T) {
	f := calculateFare(1, 0)

	if f != 4.0 {
		t.Errorf("expected 4, got %f", f)
	}
}
