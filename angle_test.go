package unit

import (
	"testing"
)

func TestAngle(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptoradian).Yoctoradians())
	assertFloatEqual(t, 1e3, (1 * Attoradian).Zeptoradians())
	assertFloatEqual(t, 1e3, (1 * Femtoradian).Attoradians())
	assertFloatEqual(t, 1e3, (1 * Picoradian).Femtoradians())
	assertFloatEqual(t, 1e3, (1 * Nanoradian).Picoradians())
	assertFloatEqual(t, 1e3, (1 * Microradian).Nanoradians())
	assertFloatEqual(t, 1e3, (1 * Milliradian).Microradians())

	assertFloatEqual(t, 1e3, (1 * Radian).Milliradians())
	assertFloatEqual(t, 1e2, (1 * Radian).Centiradians())
	assertFloatEqual(t, 1e1, (1 * Radian).Deciradians())
	assertFloatEqual(t, 1e0, (1 * Radian).Radians())

	// constant ~57.2958 from https://en.wikipedia.org/wiki/Radian#Conversion_between_radians_and_degrees
	assertFloatEqual(t, 57.295779513082321, (1 * Radian).Degrees())

	// constants from https://en.wikipedia.org/wiki/Minute_and_second_of_arc#Symbols_and_abbreviations
	assertFloatEqual(t, 17.453292519943295, (1 * Degree).Milliradians())
	assertFloatEqual(t, 290.8882086657216, (1 * Arcminute).Microradians())
	assertFloatEqual(t, 4.84813681109536, (1 * Arcsecond).Microradians())

	assertFloatEqual(t, 60.0, (1 * Degree).Arcminutes())
	assertFloatEqual(t, 60.0, (1 * Arcminute).Arcseconds())
	assertFloatEqual(t, 1e3, (1 * Arcsecond).Milliarcseconds())
	assertFloatEqual(t, 1e6, (1 * Arcsecond).Microarcseconds())
}
