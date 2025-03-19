package lab2

import (
	"errors"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	return errors.New("Some important error")
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	result, err := ParsePrefixToLisp(string(data))
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(result))
	return err
}
