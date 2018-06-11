package unit

import (
	"testing"
)

func TestAmountOfSubstance(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e0, (1 * Mole).Moles())
}
