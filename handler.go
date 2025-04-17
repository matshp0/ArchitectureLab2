package lab2

import (
	"fmt"
	"io"
)

type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Reader)
	if err != nil {
		return err
	}
	result, err := CalculatePostfix(string(data))
	if err != nil {
		return err
	}
	str := fmt.Sprintf("%g", result)
	_, err = fmt.Fprintln(ch.Writer, "Calculated postfix expression =", str)
	return nil
}
