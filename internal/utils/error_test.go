package utils_test

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/utils"
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

func TestPrintCapitalizedError(t *testing.T) {
	err := errors.New("something went wrong")

	output := captureOutput(func() {
		utils.PrintCapitalizedError(err)
	})

	expected := "Something went wrong\n"

	if output != expected {
		t.Fatalf("expected output %q, got %q", expected, output)
	}
}
