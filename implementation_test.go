package lab2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	input  string
	result float64
	isErr  bool
}

func TestCalculatePostfix(t *testing.T) {
	testCases := []testCase{
		{"2 3 +", 5.0, false},
		{"3 2 3 4 5", 0.0, true},
		{"2 a 3 2", 0.0, true},
		{"     ", 0.0, true},
		{"", 0.0, true},
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
			value, err := CalculatePostfix(test.input)
			if test.isErr {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
				assert.Equal(t, test.result, value, "Expected result does not match")
			}
		})
	}
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("2 2 + 3 3 * +")
	fmt.Println(res)
	// Output:
	// 13
}
