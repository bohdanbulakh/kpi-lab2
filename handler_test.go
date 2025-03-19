package lab2

import (
	"bytes"
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{name: "empty input", input: "", expected: "", hasError: true},
		{name: "invalid expression", input: "+ 5 * 3 a, + 1, & 3 4", expected: "", hasError: true},
		{name: "empty", input: "+ 1 2", expected: "(+ 1 2)", hasError: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			inputReader := strings.NewReader(tc.input)
			outputBuffer := &bytes.Buffer{}

			handler := ComputeHandler{
				Input:  inputReader,
				Output: outputBuffer,
			}

			err := handler.Compute()
			if err != nil && !tc.hasError {
				t.Fatalf("Compute() returned an error: %v", err)
			}

			actualOutput := outputBuffer.String()
			if actualOutput != tc.expected {
				t.Errorf("Expected output %q, but got %q", tc.expected, actualOutput)
			}
		})
	}
}
