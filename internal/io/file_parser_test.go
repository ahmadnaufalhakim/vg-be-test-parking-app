package io_test

import (
	"os"
	"testing"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/io"
)

func TestExtractCommandsFromFile_Success(t *testing.T) {
	// create a temporary file
	tmp, err := os.CreateTemp("", "commands-*.txt")
	if err != nil {
		t.Fatalf("failed to creaate temp file: %v", err)
	}
	defer os.Remove(tmp.Name()) // clean up

	content := "create_parking_lot 4\npark CAR-01\nstatus\npark CAR-02\nleave CAR-03 5"
	if _, err := tmp.Write([]byte(content)); err != nil {
		t.Fatalf("unexpected error writing to temp file: %v", err)
	}

	// close the file so the function can read the file
	tmp.Close()

	commands, err := io.ExtractCommandsFromFile(tmp.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"create_parking_lot 4",
		"park CAR-01",
		"status",
		"park CAR-02",
		"leave CAR-03 5",
	}

	if len(commands) != len(expected) {
		t.Fatalf("expected %d commands, got %d", len(expected), len(commands))
	}

	for i, cmd := range commands {
		if cmd != expected[i] {
			t.Fatalf("expected command %q at index %d, got %q", expected[i], i, cmd)
		}
	}
}

func TestExtractCommandsFromFile_FileNotFound(t *testing.T) {
	_, err := io.ExtractCommandsFromFile("this-file-does-not-exist.txt")
	if err == nil {
		t.Fatalf("expected an error when file does not exist")
	}
}

func TestExtractCommandsFromFile_EmptyFile(t *testing.T) {
	tmp, err := os.CreateTemp("", "empty-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmp.Name())

	tmp.Close()

	commands, err := io.ExtractCommandsFromFile(tmp.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(commands) != 0 {
		t.Fatalf("expected 0 commands from empty file, got %d", len(commands))
	}
}
