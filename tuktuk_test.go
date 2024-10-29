package main

import (
	"reflect"
	"testing"
)

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

func TestMinimumFare(t *testing.T) {
	t.Run("should return minimum fare as 35 THB when the fare less than 35 THB", func(t *testing.T) {
		fare := 34.0

		got := Minimum(fare)

		want := 35.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should keep the fare as 35 THB when the fare is 35 THB", func(t *testing.T) {
		fare := 35.0

		got := Minimum(fare)

		want := 35.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should keep the fare as it is when the fare more than 35 THB", func(t *testing.T) {
		fare := 36.0

		got := Minimum(fare)

		want := 36.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}

func TestRide(t *testing.T) {
	t.Run("should return minimum fare as 35 THB when the fare less than 35 THB", func(t *testing.T) {
		distance := 7.9
		waitingTime := 179

		got := Ride(distance, waitingTime)

		want := 35.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})

	t.Run("should return fare of 10km with 3m (180 seconds) waiting time", func(t *testing.T) {
		distance := 10.0
		waitingTime := 180

		got := Ride(distance, waitingTime)

		want := 43.0
		if got != want {
			t.Errorf("expected %f, got %f", want, got)
		}
	})
}

func TestMultipleRides(t *testing.T) {
	t.Run("should return fare of 10km with 3m (180 seconds) waiting time", func(t *testing.T) {
		rds := []ride{
			{distance: 7.9, waitingTime: 120},  // 8.0km, 2.0m  = 35
			{distance: 9.5, waitingTime: 180},  // 9.5km, 3.0m = 41
			{distance: 10.0, waitingTime: 299}, // 10.0km, 5.0m = 45
		}

		total := TotalFare(rds)

		want := 121.0
		if total != want {
			t.Errorf("expected %f, got %f", want, total)
		}
	})
}

func TestInvoiceDetails(t *testing.T) {
	t.Run("invoice details for 3 rides", func(t *testing.T) {
		rds := []ride{
			{distance: 12.0, waitingTime: 180}, // 4*12.0km + 3.0m = 51.0
			{distance: 7.5, waitingTime: 180},  // 4*7.5km + 3.0m = 33.0 will be 35.0
			{distance: 8.5, waitingTime: 180},  // 4*8.5km + 3.0m = 37.0
		}
		want := invoice{
			Rides:              rds,
			Total:              123.0,
			AverageFarePerRide: 41.00,
		}

		got := InvoiceDetails(rds)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v, got %v", want, got)
		}
	})
}
