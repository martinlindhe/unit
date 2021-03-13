package unit

import (
	"testing"
)

func TestPressure(t *testing.T) {

	// SI derived
	assertFloatEqual(t, 1e3, (1 * Zeptopascal).Yoctopascals())
	assertFloatEqual(t, 1e3, (1 * Attopascal).Zeptopascals())
	assertFloatEqual(t, 1e3, (1 * Femtopascal).Attopascals())
	assertFloatEqual(t, 1e3, (1 * Picopascal).Femtopascals())
	assertFloatEqual(t, 1e3, (1 * Nanopascal).Picopascals())
	assertFloatEqual(t, 1e3, (1 * Micropascal).Nanopascals())
	assertFloatEqual(t, 1e3, (1 * Millipascal).Micropascals())

	assertFloatEqual(t, 1e3, (1 * Pascal).Millipascals())
	assertFloatEqual(t, 1e2, (1 * Pascal).Centipascals())
	assertFloatEqual(t, 1e1, (1 * Pascal).Decipascals())
	assertFloatEqual(t, 1e0, (1 * Pascal).Pascals())
	assertFloatEqual(t, 1e-1, (1 * Pascal).Decapascals())
	assertFloatEqual(t, 1e-2, (1 * Pascal).Hectopascals())
	assertFloatEqual(t, 1e-3, (1 * Pascal).Kilopascals())

	assertFloatEqual(t, 1e-3, (1 * Kilopascal).Megapascals())
	assertFloatEqual(t, 1e-3, (1 * Megapascal).Gigapascals())
	assertFloatEqual(t, 1e-3, (1 * Gigapascal).Terapascals())
	assertFloatEqual(t, 1e-3, (1 * Terapascal).Petapascals())
	assertFloatEqual(t, 1e-3, (1 * Petapascal).Exapascals())
	assertFloatEqual(t, 1e-3, (1 * Exapascal).Zettapascals())
	assertFloatEqual(t, 1e-3, (1 * Zettapascal).Yottapascals())

	// non-SI
	assertFloatEqual(t, 1e3, (1 * Zeptobar).Yoctobars())
	assertFloatEqual(t, 1e3, (1 * Attobar).Zeptobars())
	assertFloatEqual(t, 1e3, (1 * Femtobar).Attobars())
	assertFloatEqual(t, 1e3, (1 * Picobar).Femtobars())
	assertFloatEqual(t, 1e3, (1 * Nanobar).Picobars())
	assertFloatEqual(t, 1e3, (1 * Microbar).Nanobars())
	assertFloatEqual(t, 1e3, (1 * Millibar).Microbars())

	assertFloatEqual(t, 1e3, (1 * Bar).Millibars())
	assertFloatEqual(t, 1e2, (1 * Bar).Centibars())
	assertFloatEqual(t, 1e1, (1 * Bar).Decibars())
	assertFloatEqual(t, 1e0, (1 * Bar).Bars())
	assertFloatEqual(t, 1e-1, (1 * Bar).Decabars())
	assertFloatEqual(t, 1e-2, (1 * Bar).Hectobars())
	assertFloatEqual(t, 1e-3, (1 * Bar).Kilobars())

	assertFloatEqual(t, 1e-3, (1 * Kilobar).Megabars())
	assertFloatEqual(t, 1e-3, (1 * Megabar).Gigabars())
	assertFloatEqual(t, 1e-3, (1 * Gigabar).Terabars())
	assertFloatEqual(t, 1e-3, (1 * Terabar).Petabars())
	assertFloatEqual(t, 1e-3, (1 * Petabar).Exabars())
	assertFloatEqual(t, 1e-3, (1 * Exabar).Zettabars())
	assertFloatEqual(t, 1e-3, (1 * Zettabar).Yottabars())

	// misc
	assertFloatEqual(t, 0.09869232667160129, (10000 * Pascal).Atmospheres())
	assertFloatEqual(t, 0.10197162129779283, (10000 * Pascal).TechAtmospheres())
	assertFloatEqual(t, 75.00615050434136, (10000 * Pascal).Torrs())
	assertFloatEqual(t, 1.4503683935719673, (10000 * Pascal).PoundsPerSquareInch())
	assertFloatEqual(t, 2.9529980164712323, (10000 * Pascal).InchOfMercury())
}
