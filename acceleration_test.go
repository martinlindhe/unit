package unit

import (
	"testing"
)

func TestAcceleration(t *testing.T) {

	// constants from https://en.wikipedia.org/wiki/Standard_gravity#Conversions
	assertFloatEqual(t, 0.03280839895013123, (1 * Gal).FeetPerSecondSquared())
	assertFloatEqual(t, 0.01, (1 * Gal).MetersPerSecondSquared())
	assertFloatEqual(t, 1e0, (1 * Gal).CentimetersPerSecondSquared()) // alias
	assertFloatEqual(t, 0.0010197162129779284, (1 * Gal).StandardGravity())

	assertFloatEqual(t, 30.4800, (1 * FootPerSecondSquared).Gals())
	assertFloatEqual(t, 0.304800, (1 * FootPerSecondSquared).MetersPerSecondSquared())
	assertFloatEqual(t, 0.031080950171567256, (1 * FootPerSecondSquared).StandardGravity())

	assertFloatEqual(t, 100.0, (1 * MeterPerSecondSquared).Gals())
	assertFloatEqual(t, 3.280839895013123, (1 * MeterPerSecondSquared).FeetPerSecondSquared())
	assertFloatEqual(t, 0.10197162129779283, (1 * MeterPerSecondSquared).StandardGravity())

	assertFloatEqual(t, 980.665, (1 * StandardGravity).Gals())
	assertFloatEqual(t, 32.17404855643044, (1 * StandardGravity).FeetPerSecondSquared())
	assertFloatEqual(t, 9.80665, (1 * StandardGravity).MetersPerSecondSquared())
}
