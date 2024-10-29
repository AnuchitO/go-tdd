package main

import "testing"

func TestCalculateFare(t *testing.T) {
	t.Run("should return fare of 1 km without waiting time", func(t *testing.T) {
		distance := 1.0
		waitingTime := 0.0

		got := calculateFare(distance, waitingTime)

		want := 4.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should return fare of 2 km without waiting time", func(t *testing.T) {
		distance := 2.0
		waitingTime := 0.0

		got := calculateFare(distance, waitingTime)

		want := 8.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}
