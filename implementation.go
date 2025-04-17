package lab2

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var opMap = map[string]func(float64, float64) (float64, error){
	"+": func(a, b float64) (float64, error) { return a + b, nil },
	"-": func(a, b float64) (float64, error) { return a - b, nil },
	"*": func(a, b float64) (float64, error) { return a * b, nil },
	"/": func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	},
	"^": func(a, b float64) (float64, error) { return math.Pow(a, b), nil },
}

type Operand struct {
	number float64
	symbol string
}

func ParseOperands(str string) ([]Operand, error) {
	var result []Operand
	arr := strings.Split(str, " ")
	for _, el := range arr {
		var parsed Operand
		value, err := strconv.ParseFloat(el, 64)
		if err != nil {
			_, ok := opMap[el]
			if ok != true {
				return nil, errors.New("invalid operator")
			}
			parsed.symbol = el
		} else {
			parsed.number = value
		}
		result = append(result, parsed)
	}
	return result, nil
}

func CalculatePostfix(input string) (float64, error) {
	operands, err := ParseOperands(input)
	if err != nil {
		return 0, err
	}
	i := 0
	for len(operands) != 1 {
		if operands[i].symbol == "" {
			i++
			if i == len(operands) {
				return 0, errors.New("invalid expression")
			}
			continue
		}
		if i < 2 {
			return 0, errors.New("invalid expression")
		}
		v1, v2 := operands[i-2].number, operands[i-1].number
		operation := opMap[operands[i].symbol]
		value, err := operation(v1, v2)
		if err != nil {
			return 0, err
		}
		var result Operand
		result.number = value
		operands = append(operands[:i-2], append([]Operand{result}, operands[i+1:]...)...)
		i -= 2
	}
	return operands[0].number, nil
}
