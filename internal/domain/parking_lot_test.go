package domain_test

import (
	"testing"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/domain"
)

func TestNewParkingLot_InvalidCapacity(t *testing.T) {
	_, err := domain.NewParkingLot(0)
	if err == nil {
		t.Fatalf("expected error for zero capacity")
	}
}

func TestParkingLot_ParkAndRemove(t *testing.T) {
	pl, err := domain.NewParkingLot(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if pl.AvailableSlots.Len() != 3 {
		t.Fatalf("expected 3 available slots, got %d", pl.AvailableSlots.Len())
	}

	// try parking a car
	slot, err := pl.ParkCar("TEST-01")
	if err != nil {
		t.Fatalf("unexpected error parking car: %v", err)
	}
	if slot == nil {
		t.Fatalf("expected a slot allocated")
	}
	if slot.ID != 1 {
		t.Fatalf("expected slot ID 1, got %d", slot.ID)
	}
	if pl.AvailableSlots.Len() != 2 {
		t.Fatalf("expected 2 available slots after parking a car, got %d", pl.AvailableSlots.Len())
	}

	// try parking the same car again,
	// should result in an error
	_, err = pl.ParkCar("TEST-01")
	if err == nil {
		t.Fatalf("expected error when parking the same car")
	}

	// try removing a non-existent car
	_, err = pl.RemoveCar("NOT-EXIST")
	if err == nil {
		t.Fatalf("expected error when removing a non-existent car")
	}

	// try removing a parked car
	removed, err := pl.RemoveCar("TEST-01")
	if err != nil {
		t.Fatalf("unexpected error removing car: %v", err)
	}
	if removed.ID != 1 {
		t.Fatalf("expected removed slot ID 1, got %v", removed.ID)
	}
	if pl.AvailableSlots.Len() != 3 {
		t.Fatalf("expected 3 available slots after removing a car, got %d", pl.AvailableSlots.Len())
	}
}

// i wanna test is empty and is full, how?
func TestParkingLot_IsEmptyAndIsFull(t *testing.T) {
	pl, err := domain.NewParkingLot(2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// initially empty: all slots are available
	if !pl.IsEmpty() {
		t.Fatalf("expected parking lot to be empty at start")
	}
	if pl.IsFull() {
		t.Fatalf("expected parking lot NOT to be full at start")
	}

	// try parking a car
	_, err = pl.ParkCar("TEST-01")
	if err != nil {
		t.Fatalf("unexpected error parking TEST-01: %v", err)
	}

	if pl.IsEmpty() {
		t.Fatalf("expected parking lot NOT to be empty after parking 1 car")
	}
	if pl.IsFull() {
		t.Fatalf("expected parking lot NOT to be full after parking 1 car")
	}

	// try parking a second car,
	// parking lot should be full
	_, err = pl.ParkCar("TEST-02")
	if err != nil {
		t.Fatalf("unexpected error parking TEST-02: %v", err)
	}

	if pl.IsEmpty() {
		t.Fatalf("expected parking lot NOT to be empty after parking 2 cars")
	}
	if !pl.IsFull() {
		t.Fatalf("expected parking lot to be full after parking 2 cars")
	}

	// try removing one car,
	// parking lot should not be full anymore
	_, err = pl.RemoveCar("TEST-01")
	if err != nil {
		t.Fatalf("unexpected error removing TEST-01: %v", err)
	}

	if pl.IsEmpty() {
		t.Fatalf("expected parking lot NOT to be empty after removing 1 car")
	}
	if pl.IsFull() {
		t.Fatalf("expected parking lot NOT to be full after removing 1 car")
	}

	// try removing the second car,
	// parking lot should be empty
	_, err = pl.RemoveCar("TEST-02")
	if err != nil {
		t.Fatalf("unexpected error removing TEST-02: %v", err)
	}

	if !pl.IsEmpty() {
		t.Fatalf("expected parking lot to be empty after removing all cars")
	}
	if pl.IsFull() {
		t.Fatalf("expected parking lot NOT to be full after removing all cars")
	}
}
