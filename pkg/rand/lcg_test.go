//go:build unit
// +build unit

package rand

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCGReset(t *testing.T) {
	lcg, err := NewLinearCongruentialGenerator(LinearCongruentialGeneratorOptions{
		Seed:       1,
		Modulus:    10,
		Multiplier: 2,
		Increment:  1,
	})

	if err != nil {
		assert.FailNow(t, fmt.Sprintf("Expected nil error return from Init, got %v", err))
	}

	first := lcg.Next()

	for i := 0; i < 4; i++ {
		lcg.Next()
	}

	lcg.Reset()

	afterReset := lcg.Next()

	assert.Equal(t, first, afterReset)
}
