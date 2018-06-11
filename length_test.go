package unit

import (
	"testing"
)

func TestLength(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptometer).Yoctometers())
	assertFloatEqual(t, 1e3, (1 * Attometer).Zeptometers())
	assertFloatEqual(t, 1e3, (1 * Femtometer).Attometers())
	assertFloatEqual(t, 1e3, (1 * Picometer).Femtometers())
	assertFloatEqual(t, 1e3, (1 * Nanometer).Picometers())
	assertFloatEqual(t, 1e3, (1 * Micrometer).Nanometers())
	assertFloatEqual(t, 1e3, (1 * Millimeter).Micrometers())

	assertFloatEqual(t, 1e3, (1 * Meter).Millimeters())
	assertFloatEqual(t, 1e2, (1 * Meter).Centimeters())
	assertFloatEqual(t, 1e1, (1 * Meter).Decimeters())
	assertFloatEqual(t, 1e0, (1 * Meter).Meters())
	assertFloatEqual(t, 1e-1, (1 * Meter).Decameters())
	assertFloatEqual(t, 1e-2, (1 * Meter).Hectometers())
	assertFloatEqual(t, 1e-3, (1 * Meter).Kilometers())
	assertFloatEqual(t, 1e-4, (1 * Meter).ScandinavianMiles())

	assertFloatEqual(t, 1e-3, (1 * Kilometer).Megameters())
	assertFloatEqual(t, 1e-3, (1 * Megameter).Gigameters())
	assertFloatEqual(t, 1e-3, (1 * Gigameter).Terameters())
	assertFloatEqual(t, 1e-3, (1 * Terameter).Petameters())
	assertFloatEqual(t, 1e-3, (1 * Petameter).Exameters())
	assertFloatEqual(t, 1e-3, (1 * Exameter).Zettameters())
	assertFloatEqual(t, 1e-3, (1 * Zettameter).Yottameters())

	// US
	assertFloatEqual(t, 0.3048, (1 * Foot).Meters())
	assertFloatEqual(t, 12, (1 * Foot).Inches())
	assertFloatEqual(t, 3, (1 * Foot).Hands())
	assertFloatEqual(t, 0.08333333333333334, (1 * Inch).Feet())
	assertFloatEqual(t, 0.3333333333333333, (1 * Foot).Yards())
	assertFloatEqual(t, 7.92, (1 * Link).Inches())
	assertFloatEqual(t, 101.6, (1 * Hand).Millimeters())
	assertFloatEqual(t, 0.18181818181818182, (1 * Yard).Rods())
	assertFloatEqual(t, 0.25, (1 * Rod).Chains())
	assertFloatEqual(t, 25., (1 * Rod).Links())
	assertFloatEqual(t, 0.1, (1 * Chain).Furlongs())
	assertFloatEqual(t, 12.5, (100 * Furlong).Miles())
	assertFloatEqual(t, 54.6806649168854, (100 * Meter).Fathoms())
	assertFloatEqual(t, 0.5399568034557236, (100 * Meter).Cables())

	// US maritime
	assertFloatEqual(t, 1.8288, (1 * Fathom).Meters())
	assertFloatEqual(t, 185.2, (1 * Cable).Meters())
	assertFloatEqual(t, 0.8689762419006479, (1 * Mile).NauticalMiles())

	// space
	assertFloatEqual(t, 389.17240036420395, (1 * AstronomicalUnit).LunarDistances())
	assertFloatEqual(t, 0.0025695552897999903, (1 * LunarDistance).AstronomicalUnits())
	assertFloatEqual(t, 63241.07708426628, (1 * LightYear).AstronomicalUnits())
}
