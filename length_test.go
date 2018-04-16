package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {

	// SI
	assert.Equal(t, 1e3, (1 * Zeptometer).Yoctometers())
	assert.Equal(t, 1000.0000000000001, (1 * Attometer).Zeptometers()) // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtometer).Attometers())
	assert.Equal(t, 999.9999999999999, (1 * Picometer).Femtometers())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Nanometer).Picometers())   // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Micrometer).Nanometers())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Millimeter).Micrometers()) // round error, expected 1e3

	assert.Equal(t, 1e3, (1 * Meter).Millimeters())
	assert.Equal(t, 1e2, (1 * Meter).Centimeters())
	assert.Equal(t, 1e1, (1 * Meter).Decimeters())
	assert.Equal(t, 1e0, (1 * Meter).Meters())
	assert.Equal(t, 1e-1, (1 * Meter).Decameters())
	assert.Equal(t, 1e-2, (1 * Meter).Hectometers())
	assert.Equal(t, 1e-3, (1 * Meter).Kilometers())
	assert.Equal(t, 1e-4, (1 * Meter).ScandinavianMiles())

	assert.Equal(t, 1e-3, (1 * Kilometer).Megameters())
	assert.Equal(t, 1e-3, (1 * Megameter).Gigameters())
	assert.Equal(t, 1e-3, (1 * Gigameter).Terameters())
	assert.Equal(t, 1e-3, (1 * Terameter).Petameters())
	assert.Equal(t, 1e-3, (1 * Petameter).Exameters())
	assert.Equal(t, 1e-3, (1 * Exameter).Zettameters())
	assert.Equal(t, 1e-3, (1 * Zettameter).Yottameters())

	// US
	assert.Equal(t, 0.30479999999999996, (1 * Foot).Meters())
	assert.Equal(t, 11.999999999999998, (1 * Foot).Inches()) // round error, expected 12
	assert.Equal(t, 2.9999999999999996, (1 * Foot).Hands())
	assert.Equal(t, 0.08333333333333334, (1 * Inch).Feet())
	assert.Equal(t, 0.3333333333333333, (1 * Foot).Yards())
	assert.Equal(t, 7.92, (1 * Link).Inches())
	assert.Equal(t, 101.6, (1 * Hand).Millimeters())
	assert.Equal(t, 0.18181818181818182, (1 * Yard).Rods())
	assert.Equal(t, 0.25, (1 * Rod).Chains())
	assert.Equal(t, 25., (1 * Rod).Links())
	assert.Equal(t, 0.1, (1 * Chain).Furlongs())
	assert.Equal(t, 12.5, (100 * Furlong).Miles())
	assert.Equal(t, 54.6806649168854, (100 * Meter).Fathoms())
	assert.Equal(t, 0.5399568034557236, (100 * Meter).Cables())

	// US maritime
	assert.Equal(t, 1.8287999999999998, (1 * Fathom).Meters()) // round error, expected 1.8288
	assert.Equal(t, 185.2, (1 * Cable).Meters())
	assert.Equal(t, 0.8689762419006479, (1 * Mile).NauticalMiles())

	// space
	assert.Equal(t, 389.17240036420395, (1 * AstronomicalUnit).LunarDistances())
	assert.Equal(t, 0.0025695552897999903, (1 * LunarDistance).AstronomicalUnits())
	assert.Equal(t, 63241.07708426628, (1 * LightYear).AstronomicalUnits())
}
