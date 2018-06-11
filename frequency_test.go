package unit

import (
	"testing"
)

func TestFrequency(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptohertz).Yoctohertz())
	assertFloatEqual(t, 1e3, (1 * Attohertz).Zeptohertz())
	assertFloatEqual(t, 1e3, (1 * Femtohertz).Attohertz())
	assertFloatEqual(t, 1e3, (1 * Picohertz).Femtohertz())
	assertFloatEqual(t, 1e3, (1 * Nanohertz).Picohertz())
	assertFloatEqual(t, 1e3, (1 * Microhertz).Nanohertz())
	assertFloatEqual(t, 1e6, (1 * Hertz).Microhertz())

	assertFloatEqual(t, 1e3, (1 * Hertz).Millihertz())
	assertFloatEqual(t, 1e2, (1 * Hertz).Centihertz())
	assertFloatEqual(t, 1e1, (1 * Hertz).Decihertz())
	assertFloatEqual(t, 1e0, (1 * Hertz).Hertz())
	assertFloatEqual(t, 1e-1, (1 * Hertz).Decahertz())
	assertFloatEqual(t, 1e-2, (1 * Hertz).Hectohertz())
	assertFloatEqual(t, 1e-3, (1 * Hertz).Kilohertz())

	assertFloatEqual(t, 1e-6, (1 * Hertz).Megahertz())
	assertFloatEqual(t, 1e-3, (1 * Megahertz).Gigahertz())
	assertFloatEqual(t, 1e-3, (1 * Gigahertz).Terahertz())
	assertFloatEqual(t, 1e-3, (1 * Terahertz).Petahertz())
	assertFloatEqual(t, 1e-3, (1 * Petahertz).Exahertz())
	assertFloatEqual(t, 1e-3, (1 * Exahertz).Zettahertz())
	assertFloatEqual(t, 1e-3, (1 * Zettahertz).Yottahertz())
}
