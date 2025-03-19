package lab2

import (
	"fmt"
	"testing"
)

func TestParsePrefixToLisp(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{name: "addition", input: "+ 1 2", expected: "(+ 1 2)", hasError: false},
		{name: "multiplication", input: "* 3 4", expected: "(* 3 4)", hasError: false},
		{name: "subtraction", input: "- 7 2", expected: "(- 7 2)", hasError: false},
		{name: "division", input: "/ 8 4", expected: "(/ 8 4)", hasError: false},
		{name: "power", input: "^ 2 3", expected: "(pow 2 3)", hasError: false},

		{name: "complex expression 1", input: "+ 5 * - 4 2 ^ 3 2", expected: "(+ 5 (* (- 4 2) (pow 3 2)))", hasError: false},
		{name: "complex expression 2", input: "- + 1 2 * 3 4", expected: "(- (+ 1 2) (* 3 4))", hasError: false},
		{name: "complex expression 3", input: "* - + 2 3 4 / 9 3", expected: "(* (- (+ 2 3) 4) (/ 9 3))", hasError: false},
		{name: "complex expression 4", input: "+ 10 - 8 * 2 4", expected: "(+ 10 (- 8 (* 2 4)))", hasError: false},
		{name: "complex expression 5", input: "/ + 20 5 - 10 2", expected: "(/ (+ 20 5) (- 10 2))", hasError: false},

		{name: "empty input", input: "", expected: "", hasError: true},
		{name: "invalid character", input: "+ 5 * 3 a", expected: "", hasError: true},
		{name: "missing operand", input: "+ 1", expected: "", hasError: true},
		{name: "invalid operator", input: "& 3 4", expected: "", hasError: true},
		{name: "extra spaces", input: "  +  1  2  ", expected: "(+ 1 2)", hasError: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParsePrefixToLisp(tc.input)
			if (err != nil) != tc.hasError {
				t.Errorf("unexpected error for input %q: %v", tc.input, err)
			}
			if result != tc.expected {
				t.Errorf("for input %q expected %q but got %q", tc.input, tc.expected, result)
			}
		})
	}
}

func ExampleParsePrefixToLisp() {
	result, _ := ParsePrefixToLisp("+ 5 * - 4 2 ^ 3 2")
	fmt.Println(result)
	//Output: (+ 5 (* (- 4 2) (pow 3 2)))
}
