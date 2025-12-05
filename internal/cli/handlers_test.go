package cli_test

import (
	"testing"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/cli"
	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/testutils"
)

func TestCreateParkingLot_Success(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.CreateParkingLot([]string{"5"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if d.GetParkingLot() == nil {
		t.Fatalf("parking lot should be initialized")
	}
}

func TestCreateParkingLot_InvalidArgs(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.CreateParkingLot([]string{})
	if err == nil {
		t.Fatalf("expected error for missing args")
	}

	err = d.CreateParkingLot([]string{"not-int"})
	if err == nil {
		t.Fatalf("expected error for invalid capacity string")
	}
}

func TestCreateParkingLot_InvalidCapacity(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.CreateParkingLot([]string{"-3"})
	if err == nil {
		t.Fatalf("expected error for invalid capacity")
	}
}

func TestPark_Success(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.CreateParkingLot([]string{"2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := testutils.CaptureOutput(func() {
		err := d.Park([]string{"KA-01-HH-1234"})
		if err != nil {
			t.Fatalf("unexpected park error: %v", err)
		}
	})
	expected := "Allocated slot number: 1\n"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestPark_InvalidArgs(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Park([]string{})
	if err == nil {
		t.Fatalf("expected error for wrong arg count")
	}
}

func TestPark_NoParkingLot(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Park([]string{"KA-01-HH-1234"})
	if err == nil {
		t.Fatalf("expected error for uninitialized parking lot")
	}
}

func TestLeave_Success(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.CreateParkingLot([]string{"2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = d.Park([]string{"KA-01-HH-1234"})
	if err != nil {
		t.Fatalf("unexpected park error: %v", err)
	}

	out := testutils.CaptureOutput(func() {
		err := d.Leave([]string{"KA-01-HH-1234", "4"})
		if err != nil {
			t.Fatalf("unexpected leave error: %v", err)
		}
	})
	expected := "Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30\n"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestLeave_InvalidArgs(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Leave([]string{"KA-01-HH-1234"})
	if err == nil {
		t.Fatalf("expected error for wrong arg count")
	}
}

func TestLeave_NotInitialized(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Leave([]string{"KA-01-HH-1234", "2"})
	if err == nil {
		t.Fatalf("expected error when parking lot not initialized")
	}
}

func TestLeave_InvalidHours(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.CreateParkingLot([]string{"2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = d.Park([]string{"KA-01-HH-1234"})
	if err != nil {
		t.Fatalf("unexpected park error: %v", err)
	}

	err = d.Leave([]string{"KA-01-HH-1234", "-4"})
	if err == nil {
		t.Fatalf("expected error for invalid hours")
	}

	err = d.Leave([]string{"KA-01-HH-1234", "not-number"})
	if err == nil {
		t.Fatalf("expected error for invalid hours format")
	}
}

func TestStatus_Success(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.CreateParkingLot([]string{"2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = d.Park([]string{"KA-01-HH-1234"})
	if err != nil {
		t.Fatalf("unexpected park error: %v", err)
	}

	err = d.Park([]string{"KA-01-HH-9999"})
	if err != nil {
		t.Fatalf("unexpected park error: %v", err)
	}

	out := testutils.CaptureOutput(func() {
		err := d.Status([]string{})
		if err != nil {
			t.Fatalf("unexpected status error: %v", err)
		}
	})
	expected :=
		"Slot No.\tRegistration No.\n" +
			"1\tKA-01-HH-1234\n" +
			"2\tKA-01-HH-9999\n"
	if out != expected {
		t.Fatalf("\nEXPECTED:\n%s\nGOT:\n%s", expected, out)
	}
}

func TestStatus_NotInitialized(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Status([]string{})
	if err == nil {
		t.Fatalf("expected error: parking lot not initialized")
	}
}
