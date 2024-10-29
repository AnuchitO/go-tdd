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

	t.Run("should return fare of 0 km with 1 minute waiting time", func(t *testing.T) {
		distance := 0.0
		waitingTime := 1.0

		got := calculateFare(distance, waitingTime)

		want := 1.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}

func TestRoundingDistance(t *testing.T) {
	t.Run("should round up distance from 0.1km to 0.5km", func(t *testing.T) {
		distance := 0.1

		got := roundUpDistance(distance)

		want := 0.5
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should keep distance be a 0.5km", func(t *testing.T) {
		distance := 0.5

		got := roundUpDistance(distance)

		want := 0.5
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}
