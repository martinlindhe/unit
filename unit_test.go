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

// TestAssertFloatEqual tests the assertFloatEqual method.
func TestAssertFloatEqual(t *testing.T) {
	// assertFloatNotEqual verifies that calling assertFloatEqual with
	// the given floats causes the test to fail.
	assertFloatNotEqual := func(a, b float64, args ...interface{}) {
		testT := &testing.T{}
		assertFloatEqual(testT, a, b)
		if !testT.Failed() {
			assert.Fail(t, "expected assertFloatEqual to fail", args...)
		}
	}

	for _, base := range []float64{
		1e-48, 1e-24, 1e-22,
		1e-9, 1e-6, 1e-3,
		0.01, 0.1, 1, 10, 100,
		1e3, 1e6, 1e9,
		1e22, 1e24, 1e48,
	} {
		// Values within 1 part in 5 quadrillion are considered equal,
		// even though they are not strictly equal.
		smallDelta := base / 5e15
		assert.NotEqual(t, base, base+smallDelta, "%g+%g", base, smallDelta)
		assertFloatEqual(t, base, base+smallDelta, "%g+%g", base, smallDelta)

		assert.NotEqual(t, base, base-smallDelta, "%g-%g", base, smallDelta)
		assertFloatEqual(t, base, base-smallDelta, "%g-%g", base, smallDelta)

		// But an error of 1 part in a quadrillion or larger should fail.
		largeDelta := base / 1e15
		assertFloatNotEqual(base, base+largeDelta, "%g+%g", base, largeDelta)
		assertFloatNotEqual(base, base-largeDelta, "%g-%g", base, largeDelta)

		// And some other obviously different values should fail as well.
		assertFloatNotEqual(base, base+base, "%g+%g", base, base)
		if base != 1 {
			assertFloatNotEqual(base, base*base, "%g*%g", base, base)
		}
		assertFloatNotEqual(base, 0, "%g != 0", base)
	}

	assertFloatNotEqual(0, 1e-321, "0 != very small value")
	assertFloatNotEqual(0, 1e255, "0 != very large value")
	assertFloatNotEqual(1e255, 1.1e255, "large value != very large value")

	// testAccumulatedRoundingError tests that accumulated rounding errors
	// from repeated multiplications can be ignored by assertFloatEqual.
	testAccumulatedRoundingError := func(factor float64, times int, expected float64) {
		var actual = 1.0
		for i := 0; i < times; i++ {
			actual *= factor
		}
		assert.NotEqual(t, expected, actual, "accumulated rounding error")
		assertFloatEqual(t, expected, actual)
	}

	testAccumulatedRoundingError(1e-2, 10, 1e-20)
	testAccumulatedRoundingError(1e-3, 11, 1e-33)

	testAccumulatedRoundingError(1e3, 11, 1e33)
	testAccumulatedRoundingError(1e4, 13, 1e52)
}
