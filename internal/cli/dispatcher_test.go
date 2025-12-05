package cli_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/cli"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}

func TestDispatcher_UnknownCommand(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Handle("does_not_exist", []string{})
	if err == nil {
		t.Fatalf("expected error for unknown command")
	}
}

func TestDispatcher_CreateParkingLot(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Handle("create_parking_lot", []string{"6"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if d.Handle("status", []string{}) == nil && d.Handle("park", []string{"ABC"}) == nil {
		// parkingLot is initialized → pass
	} else {
		t.Fatalf("parking lot should have been initialized after create_parking_lot")
	}
}

func TestDispatcher_CreateParkingLot_InvalidArgs(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Handle("create_parking_lot", []string{})
	if err == nil {
		t.Fatalf("expected error for missing arguments")
	}

	err = d.Handle("create_parking_lot", []string{"-6"})
	if err == nil {
		t.Fatalf("expected error for invalid capacity")
	}

	err = d.Handle("create_parking_lot", []string{"not-number"})
	if err == nil {
		t.Fatalf("expected error for invalid capacity format")
	}
}

func TestDispatcher_Park_Success(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Handle("create_parking_lot", []string{"2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := captureOutput(func() {
		err := d.Handle("park", []string{"KA-01-HH-1234"})
		if err != nil {
			t.Fatalf("unexpected park error: %v", err)
		}
	})

	expected := "Allocated slot number: 1\n"
	if output != expected {
		t.Fatalf("expected %q got %q", expected, output)
	}
}

func TestDispatcher_Park_NoLot(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Handle("park", []string{"KA-01-HH-1234"})
	if err == nil {
		t.Fatalf("expected error: parking lot not initialized")
	}
}

func TestDispatcher_Leave_Success(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Handle("create_parking_lot", []string{"2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = d.Handle("park", []string{"KA-01-HH-1234"})
	if err != nil {
		t.Fatalf("unexpected park error: %v", err)
	}

	output := captureOutput(func() {
		err := d.Handle("leave", []string{"KA-01-HH-1234", "4"})
		if err != nil {
			t.Fatalf("unexpected leave error: %v", err)
		}
	})

	expected := "Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30\n"
	if output != expected {
		t.Fatalf("expected %q got %q", expected, output)
	}
}

func TestDispatcher_Leave_NotFound(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Handle("create_parking_lot", []string{"2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = d.Handle("leave", []string{"NOT-FOUND", "4"})
	if err == nil {
		t.Fatalf("expected error for non-existent car")
	}
}

func TestDispatcher_Status(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Handle("create_parking_lot", []string{"2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = d.Handle("park", []string{"KA-01-HH-1234"})
	if err != nil {
		t.Fatalf("unexpected park error: %v", err)
	}

	err = d.Handle("park", []string{"KA-01-HH-9999"})
	if err != nil {
		t.Fatalf("unexpected park error: %v", err)
	}

	output := captureOutput(func() {
		err := d.Handle("status", []string{})
		if err != nil {
			t.Fatalf("unexpected status error: %v", err)
		}
	})

	expected :=
		"Slot No.\tRegistration No.\n" +
			"1\tKA-01-HH-1234\n" +
			"2\tKA-01-HH-9999\n"

	if output != expected {
		t.Fatalf("\nEXPECTED:\n%s\nGOT:\n%s", expected, output)
	}
}

func TestDispatcher_Status_NoLot(t *testing.T) {
	d := cli.NewDispatcher()

	err := d.Handle("status", []string{})
	if err == nil {
		t.Fatalf("expected error: parking lot not initialized")
	}
}
