package utils_test

import (
	"errors"
	"testing"

	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/testutils"
	"github.com/ahmadnaufalhakim/vg-be-test-parking-app/internal/utils"
)

func TestPrintCapitalizedError(t *testing.T) {
	err := errors.New("something went wrong")

	output := testutils.CaptureOutput(func() {
		utils.PrintCapitalizedError(err)
	})

	expected := "Something went wrong\n"

	if output != expected {
		t.Fatalf("expected output %q, got %q", expected, output)
	}
}
