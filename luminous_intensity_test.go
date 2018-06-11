package unit

import (
	"testing"
)

func TestLuminousIntensity(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e0, (1 * Candela).Candela())
}
