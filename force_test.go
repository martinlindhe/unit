package unit

import (
	"testing"
)

func TestForce(t *testing.T) {

	// constants from https://en.wikipedia.org/wiki/Force#Units_of_measurement
	assertFloatEqual(t, 1e5, (1 * Newton).Dynes())
	assertFloatEqual(t, 0.10197162129779283, (1 * Newton).KilogramForce())
	assertFloatEqual(t, 0.2248089430997105, (1 * Newton).PoundForce())
	assertFloatEqual(t, 7.233013851209893, (1 * Newton).Poundals())

	assertFloatEqual(t, 1e-5, (1 * Dyne).Newtons())
	assertFloatEqual(t, 1.0197162129779284e-6, (1 * Dyne).KilogramForce())
	assertFloatEqual(t, 2.2480894309971053e-6, (1 * Dyne).PoundForce())
	assertFloatEqual(t, 7.233013851209893e-05, (1 * Dyne).Poundals())

	assertFloatEqual(t, 9.80665, (1 * KilogramForce).Newtons())
	assertFloatEqual(t, 980665, (1 * KilogramForce).Dynes())
	assertFloatEqual(t, 2.2046226218487757, (1 * KilogramForce).PoundForce())
	assertFloatEqual(t, 70.9316352839675, (1 * KilogramForce).Poundals())

	assertFloatEqual(t, 4.4482216152605, (1 * PoundForce).Newtons())
	assertFloatEqual(t, 444822.16152604995, (1 * PoundForce).Dynes())
	assertFloatEqual(t, 0.45359236999999997, (1 * PoundForce).KilogramForce())
	assertFloatEqual(t, 32.174048556430435, (1 * PoundForce).Poundals())

	assertFloatEqual(t, 0.13825495437600002, (1 * Poundal).Newtons())
	assertFloatEqual(t, 13825.4954376, (1 * Poundal).Dynes())
	assertFloatEqual(t, 0.0140980818501731, (1 * Poundal).KilogramForce())
	assertFloatEqual(t, 0.03108095017156726, (1 * Poundal).PoundForce())
}

func TestPoundForceExactDefinitions(t *testing.T) {
	// 1 lbf is defined as 0.45359237 kg * 9.80665 m/s^2
	assertFloatEqual(t, 4.4482216152605, (1 * PoundForce).Newtons())
	assertFloatEqual(t, 0.45359237, (1 * PoundForce).KilogramForce())
	// 1 pdl is defined as 1 lb * 1 ft/s^2 = 0.45359237 kg * 0.3048 m/s^2
	assertFloatEqual(t, 0.138254954376, (1 * Poundal).Newtons())
}
