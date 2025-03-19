package lab2

import (
	"fmt"
	"strings"
	"unicode"
)

func ParsePrefixToLisp(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("empty input")
	}

	tokens := strings.Fields(input)

	result, nextIndex, err := parseTokens(tokens, 0)
	if err != nil {
		return "", err
	}

	if nextIndex != len(tokens) {
		return "", fmt.Errorf("extra tokens in the expression")
	}

	return result, nil
}

func parseTokens(tokens []string, index int) (string, int, error) {
	if index >= len(tokens) {
		return "", index, fmt.Errorf("unexpected end of input")
	}

	token := tokens[index]
	index++

	if isOperand(token) {
		return token, index, nil
	}

	if !isOperator(token) {
		return "", index, fmt.Errorf("invalid token: %s", token)
	}

	left, newIndex, err := parseTokens(tokens, index)
	if err != nil {
		return "", newIndex, err
	}

	right, newIndex, err := parseTokens(tokens, newIndex)
	if err != nil {
		return "", newIndex, err
	}

	if token == "^" {
		token = "pow"
	}

	return fmt.Sprintf("(%s %s %s)", token, left, right), newIndex, nil
}

func isOperator(token string) bool {
	operators := map[string]bool{"+": true, "-": true, "*": true, "/": true, "^": true}
	return operators[token]
}

func isOperand(token string) bool {
	for _, ch := range token {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}
