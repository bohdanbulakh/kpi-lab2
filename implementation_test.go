package lab2

import (
	"fmt"
	"testing"
)

func TestPrefixToPostfix(t *testing.T) {

}

func ExamplePrefixToPostfix() {
	res, _ := ParsePrefixToLisp("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
