package unit

import (
	"testing"
)

func TestVoltage(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptovolt).Yoctovolts())
	assertFloatEqual(t, 1e3, (1 * Attovolt).Zeptovolts())
	assertFloatEqual(t, 1e3, (1 * Femtovolt).Attovolts())
	assertFloatEqual(t, 1e3, (1 * Picovolt).Femtovolts())
	assertFloatEqual(t, 1e3, (1 * Nanovolt).Picovolts())
	assertFloatEqual(t, 1e3, (1 * Microvolt).Nanovolts())
	assertFloatEqual(t, 1e3, (1 * Millivolt).Microvolts())

	assertFloatEqual(t, 1e3, (1 * Volt).Millivolts())
	assertFloatEqual(t, 1e2, (1 * Volt).Centivolts())
	assertFloatEqual(t, 1e1, (1 * Volt).Decivolts())
	assertFloatEqual(t, 1e0, (1 * Volt).Volts())
	assertFloatEqual(t, 1e-1, (1 * Volt).Decavolts())
	assertFloatEqual(t, 1e-2, (1 * Volt).Hectovolts())
	assertFloatEqual(t, 1e-3, (1 * Volt).Kilovolts())

	assertFloatEqual(t, 1e-3, (1 * Kilovolt).Megavolts())
	assertFloatEqual(t, 1e-3, (1 * Megavolt).Gigavolts())
	assertFloatEqual(t, 1e-3, (1 * Gigavolt).Teravolts())
	assertFloatEqual(t, 1e-3, (1 * Teravolt).Petavolts())
	assertFloatEqual(t, 1e-3, (1 * Petavolt).Exavolts())
	assertFloatEqual(t, 1e-3, (1 * Exavolt).Zettavolts())
	assertFloatEqual(t, 1e-3, (1 * Zettavolt).Yottavolts())
}
