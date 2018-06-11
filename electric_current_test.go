package unit

import (
	"testing"
)

func TestElectricCurrent(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptoampere).Yoctoamperes())
	assertFloatEqual(t, 1e3, (1 * Attoampere).Zeptoamperes())
	assertFloatEqual(t, 1e3, (1 * Femtoampere).Attoamperes())
	assertFloatEqual(t, 1e3, (1 * Picoampere).Femtoamperes())
	assertFloatEqual(t, 1e3, (1 * Nanoampere).Picoamperes())
	assertFloatEqual(t, 1e3, (1 * Microampere).Nanoamperes())
	assertFloatEqual(t, 1e3, (1 * Milliampere).Microamperes())

	assertFloatEqual(t, 1e3, (1 * Ampere).Milliamperes())
	assertFloatEqual(t, 1e2, (1 * Ampere).Deciamperes())
	assertFloatEqual(t, 1e1, (1 * Ampere).Centiamperes())
	assertFloatEqual(t, 1e0, (1 * Ampere).Amperes())
	assertFloatEqual(t, 1e-1, (1 * Ampere).Decaamperes())
	assertFloatEqual(t, 1e-2, (1 * Ampere).Hectoamperes())
	assertFloatEqual(t, 1e-3, (1 * Ampere).Kiloamperes())

	assertFloatEqual(t, 1e-3, (1 * Kiloampere).Megaamperes())
	assertFloatEqual(t, 1e-3, (1 * Megaampere).Gigaamperes())
	assertFloatEqual(t, 1e-3, (1 * Gigaampere).Teraamperes())
	assertFloatEqual(t, 1e-3, (1 * Teraampere).Petaamperes())
	assertFloatEqual(t, 1e-3, (1 * Petaampere).Exaamperes())
	assertFloatEqual(t, 1e-3, (1 * Exaampere).Zettaamperes())
	assertFloatEqual(t, 1e-3, (1 * Zettaampere).Yottaamperes())
}
