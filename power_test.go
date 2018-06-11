package unit

import (
	"testing"
)

func TestPower(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptowatt).Yoctowatts())
	assertFloatEqual(t, 1e3, (1 * Attowatt).Zeptowatts())
	assertFloatEqual(t, 1e3, (1 * Femtowatt).Attowatts())
	assertFloatEqual(t, 1e3, (1 * Picowatt).Femtowatts())
	assertFloatEqual(t, 1e3, (1 * Nanowatt).Picowatts())
	assertFloatEqual(t, 1e3, (1 * Microwatt).Nanowatts())
	assertFloatEqual(t, 1e3, (1 * Milliwatt).Microwatts())

	assertFloatEqual(t, 1e3, (1 * Watt).Milliwatts())
	assertFloatEqual(t, 1e2, (1 * Watt).Centiwatts())
	assertFloatEqual(t, 1e1, (1 * Watt).Deciwatts())
	assertFloatEqual(t, 1e0, (1 * Watt).Watts())
	assertFloatEqual(t, 1e-1, (1 * Watt).Decawatts())
	assertFloatEqual(t, 1e-2, (1 * Watt).Hectowatts())
	assertFloatEqual(t, 1e-3, (1 * Watt).Kilowatts())

	assertFloatEqual(t, 1e-3, (1 * Kilowatt).Megawatts())
	assertFloatEqual(t, 1e-3, (1 * Megawatt).Gigawatts())
	assertFloatEqual(t, 1e-3, (1 * Gigawatt).Terawatts())
	assertFloatEqual(t, 1e-3, (1 * Terawatt).Petawatts())
	assertFloatEqual(t, 1e-3, (1 * Petawatt).Exawatts())
	assertFloatEqual(t, 1e-3, (1 * Exawatt).Zettawatts())
	assertFloatEqual(t, 1e-3, (1 * Zettawatt).Yottawatts())

	// non-SI
	assertFloatEqual(t, 13.596216173039045, (10000 * Watt).Pferdestarke())
}
