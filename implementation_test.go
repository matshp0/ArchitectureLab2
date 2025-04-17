package lab2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePostfix(t *testing.T) {
	res, err := CalculatePostfix("4 2 23 3 3 3 3 3")
	if assert.Nil(t, err) {
		assert.Equal(t, 23.0, res)
	}
}

func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
