package unit

import (
	"testing"
)

func TestIlluminance(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e0, (1 * Lux).Lux())
}
