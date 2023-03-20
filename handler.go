package lab2

import (
	"bytes"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	bufRead := make([]byte, 128)
	_, err := ch.Input.Read(bufRead)
	if err != nil {
		return err
	}
	bufRead = bytes.Trim(bufRead, "\x00")

	expression := string(bufRead)
	trimmed := strings.Trim(expression, " \n")
	res, err := PostfixToInfix(trimmed)
	if err != nil {
		return err
	}
	ch.Output.Write([]byte(res))
	return nil
}
