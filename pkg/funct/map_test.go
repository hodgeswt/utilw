//go:build unit
// +build unit

package funct

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_NoErr(t *testing.T) {
	x := []int{3, 2, 5, 7, 9}
	yExpected := []int{4, 3, 6, 8, 10}

	yActual, err := Map(x, func(a int) (int, error) {
		return a + 1, nil
	})

	assert.Nil(t, err)
	assert.Equal(t, yExpected, yActual)
}

func TestMap_NoErrChangeType(t *testing.T) {
	x := []int{3, 2, 5, 7, 9}
	yExpected := []string{"3", "2", "5", "7", "9"}

	yActual, err := Map(x, func(a int) (string, error) {
		return strconv.Itoa(a), nil
	})

	assert.Nil(t, err)
	assert.Equal(t, yExpected, yActual)
}

func TestMap_Err(t *testing.T) {
	x := []int{3, 2, 5, 7, 9}

	errTest := errors.New("ErrTest")

	yActual, err := Map(x, func(a int) (string, error) {
		return "", errTest
	})

	assert.EqualError(t, err, errTest.Error())
	assert.Nil(t, yActual)
}

