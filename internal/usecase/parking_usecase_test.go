package usecase_test

import (
	"testing"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/domain"
	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/testutils"
	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/usecase"
)

func TestCreateParkingLot(t *testing.T) {
	pl, err := usecase.CreateParkingLot(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pl == nil {
		t.Fatalf("expected non-nil parking lot")
	}
	if pl.Capacity != 3 {
		t.Fatalf("expected capacity 3, got %d", pl.Capacity)
	}
}

func TestCreateParkingLot_InvalidCapacity(t *testing.T) {
	_, err := usecase.CreateParkingLot(0)
	if err == nil {
		t.Fatalf("expected error for zero capacity")
	}

	_, err = usecase.CreateParkingLot(-5)
	if err == nil {
		t.Fatalf("expected error for negative capacity")
	}
}

func TestParkCar(t *testing.T) {
	pl, _ := domain.NewParkingLot(2)

	slot, err := usecase.ParkCar(pl, "KA-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if slot.ID != 1 {
		t.Fatalf("expected slot 1, got %d", slot.ID)
	}

	// parking the same car should return error
	_, err = usecase.ParkCar(pl, "KA-01")
	if err == nil {
		t.Fatalf("expected error for duplicate parking")
	}
}

func TestRemoveCar(t *testing.T) {
	pl, _ := domain.NewParkingLot(2)

	_, err := usecase.ParkCar(pl, "KA-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	slot, err := usecase.RemoveCar(pl, "KA-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if slot.ID != 1 {
		t.Fatalf("expected removed slot ID 1, got %d", slot.ID)
	}

	// Removing non-existent should error
	_, err = usecase.RemoveCar(pl, "NOTEXIST")
	if err == nil {
		t.Fatalf("expected error for non-existent car")
	}
}

func TestCalculateParkingCharge(t *testing.T) {
	tests := []struct {
		hours    int
		expected int
		hasErr   bool
	}{
		{-1, 0, true},
		{0, 10, false},
		{1, 10, false},
		{2, 10, false}, // 10 * (2-1)
		{3, 20, false}, // 10 * (3-1)
	}

	for _, tt := range tests {
		charge, err := usecase.CalculateParkingCharge(tt.hours)

		if tt.hasErr && err == nil {
			t.Fatalf("expected error for hours=%d", tt.hours)
		}
		if !tt.hasErr && err != nil {
			t.Fatalf("unexpected error for hours=%d: %v", tt.hours, err)
		}
		if charge != tt.expected {
			t.Fatalf("expected %d, got %d for hours=%d",
				tt.expected, charge, tt.hours)
		}
	}
}

func TestStatus_EmptyLot(t *testing.T) {
	pl, _ := domain.NewParkingLot(3)

	output := testutils.CaptureOutput(func() {
		err := usecase.Status(pl)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	expected := "Slot No.\tRegistration No.\n"
	if output != expected {
		t.Fatalf("expected:\n%q\ngot:\n%q", expected, output)
	}
}

func TestStatus_TwoCars_1And2(t *testing.T) {
	pl, _ := domain.NewParkingLot(3)

	// occupy slot 1
	_, err := usecase.ParkCar(pl, "CAR-A")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// occupy slot 2
	_, err = usecase.ParkCar(pl, "CAR-B")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := testutils.CaptureOutput(func() {
		err := usecase.Status(pl)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	expected :=
		"Slot No.\tRegistration No.\n" +
			"1\tCAR-A\n" +
			"2\tCAR-B\n"
	if output != expected {
		t.Fatalf("expected:\n%q\ngot:\n%q", expected, output)
	}
}

func TestStatus_TwoCars_1And3(t *testing.T) {
	pl, _ := domain.NewParkingLot(3)

	// occupy slot 1 with CAR-A
	_, err := usecase.ParkCar(pl, "CAR-A")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// occupy slot 2 with CAR-B
	_, err = usecase.ParkCar(pl, "CAR-B")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// occupy slot 3 with CAR-C
	_, err = usecase.ParkCar(pl, "CAR-C")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// CAR-B in slot 2 leaves
	_, err = usecase.RemoveCar(pl, "CAR-B")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := testutils.CaptureOutput(func() {
		err := usecase.Status(pl)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	expected :=
		"Slot No.\tRegistration No.\n" +
			"1\tCAR-A\n" +
			"3\tCAR-C\n"
	if output != expected {
		t.Fatalf("expected:\n%q\ngot:\n%q", expected, output)
	}
}
