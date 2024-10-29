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

	t.Run("should round up distance from 0.6km to 1.0km", func(t *testing.T) {
		distance := 0.6

		got := roundUpDistance(distance)

		want := 1.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should round up distance from 1.3km to 1.5km", func(t *testing.T) {
		distance := 1.3

		got := roundUpDistance(distance)

		want := 1.5

		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}

func TestRoundingWaitingTime(t *testing.T) {
	t.Run("should round up waiting time from 29s to 1m", func(t *testing.T) {
		waitingTime := 29

		got := roundUpWaitingTime(waitingTime)

		want := 1.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should keep waiting to be 1m", func(t *testing.T) {
		waitingTime := 60

		got := roundUpWaitingTime(waitingTime)

		want := 1.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should round up waiting time from 61s to 2m", func(t *testing.T) {
		waitingTime := 61

		got := roundUpWaitingTime(waitingTime)

		want := 2.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}

func TestFare(t *testing.T) {
	t.Run("should return fare of 1.3km without waiting time", func(t *testing.T) {
		distance := 1.3
		waitingTime := 0

		got := Fare(distance, waitingTime)

		want := 6.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should return fare of 1.5km without waiting time", func(t *testing.T) {
		distance := 1.5
		waitingTime := 0

		got := Fare(distance, waitingTime)

		want := 6.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should return fare of 2km without waiting time", func(t *testing.T) {
		distance := 2.0
		waitingTime := 0

		got := Fare(distance, waitingTime)

		want := 8.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should return fare of 0km with 2m (90 seconds) waiting time", func(t *testing.T) {
		distance := 0.0
		waitingTime := 90

		got := Fare(distance, waitingTime)

		want := 2.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should return fare of 7.6km with 8m (450 seconds) waiting time", func(t *testing.T) {
		distance := 7.6
		waitingTime := 450

		got := Fare(distance, waitingTime)

		want := 40.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}
