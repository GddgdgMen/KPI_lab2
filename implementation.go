package lab2

import (
	"fmt"
	"strings"
)

func PostfixToInfix(input string) (string, error) {
	var stack []string
	const ops = "+-*/^"
	var lit = ""
	var e1 = ""
	var e2 = ""
	var res = ""

	if len(input) == 0 {
		return "", ThrowError()
	}

	var arrInput = strings.Split(input, " ")

	for i := 0; i <= len(arrInput)-1; i++ {
		lit = arrInput[i]

		if !strings.Contains(ops, lit) {
			stack = append(stack, lit)
			continue
		}
		if len(stack) < 2 {
			return "", ThrowError()
		}

		e1, stack = Pop(stack)
		e2, stack = Pop(stack)

		e1 = SetElement(e1, lit)
		e2 = SetElement(e2, lit)

		var new = e2 + " " + lit + " " + e1
		stack = append(stack, new)
	}

	res, stack = Pop(stack)

	if len(stack) > 0 {
		return "", ThrowError()
	}

	return res, nil

}

func SetElement(e string, literal string) string {
	if strings.Contains(e, " ") && strings.Contains("*/^", literal) {
		e = "(" + e + ")"
	}
	return e
}

func Pop(array []string) (string, []string) {
	if len(array) == 0 {
		return "", []string{}
	}
	return array[len(array)-1], array[:len(array)-1]
}

func ThrowError() error {
	return fmt.Errorf("invalid input")
}
