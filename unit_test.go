package unit

import (
	"math"
	"testing"

	"github.com/stretchrcom/testify/assert"
)

// assertFloatEqual asserts that the actual float64 value is within
// two epsilons of the expected float64 value.
func assertFloatEqual(t *testing.T, expected, actual float64, args ...interface{}) {
	epsilon := math.Abs(math.Nextafter(expected, actual) - expected)
	epsilon += math.Abs(math.Nextafter(actual, expected) - actual)
	assert.InDelta(t, expected, actual, epsilon, args...)
}
