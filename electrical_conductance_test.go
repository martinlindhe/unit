package unit

import (
	"testing"
)

func TestElectricalConductance(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e0, (1 * Siemens).Siemens())
}
