package lab2

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func format(f float64) string {
	return fmt.Sprintf("Calculated postfix expression = %g\n", f)
}

func TestHandler(t *testing.T) {
	testCases := []testCase{
		{"a, b, c", 5.0, true},
		{"1 2 * 2 2 a", 0.0, true},
		{"", 0.0, true},
		{"     ", 0.0, true},
		{"5 1 2 + 4 * + 3 -", 14.0, false},
		{"4 2 3 5 1 - + * +", 18.0, false},
		{"4 2 - 3 2 ^ * 5 +", 23.0, false},
	}

	res, err := CalculatePostfix("2 2 2 2 + + +")
	if assert.Nil(t, err) {
		assert.Equal(t, 8.0, res)
	}
	for _, test := range testCases {
		t.Run(test.input, func(t *testing.T) {
			var buf bytes.Buffer
			handler := ComputeHandler{}
			handler.Reader = strings.NewReader(test.input)
			handler.Writer = &buf
			err := handler.Compute()
			if test.isErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
				assert.Equal(t, format(test.result), buf.String(), "Expected result does not match")
			}
		})
	}
}
