package unit

import (
	"testing"
)

func TestElectricalResistance(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e0, (1 * Ohm).Ohms())
}
