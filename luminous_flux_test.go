package unit

import (
	"testing"
)

func TestLuminousFlux(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e0, (1 * Lumen).Lumen())
}
