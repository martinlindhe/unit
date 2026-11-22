package unit

import (
	"testing"
)

func TestForce(t *testing.T) {

	// constants from https://en.wikipedia.org/wiki/Force#Units_of_measurement
	assertFloatEqual(t, 1e5, (1 * Newton).Dynes())
	assertFloatEqual(t, 0.10197162129779283, (1 * Newton).KilogramForce())
	assertFloatEqual(t, 0.22480894309971047, (1 * Newton).PoundForce())
	assertFloatEqual(t, 7.233013851209894, (1 * Newton).Poundals())

	assertFloatEqual(t, 1e-5, (1 * Dyne).Newtons())
	assertFloatEqual(t, 1.0197162129779284e-6, (1 * Dyne).KilogramForce())
	assertFloatEqual(t, 2.248089430997105e-06, (1 * Dyne).PoundForce())
	assertFloatEqual(t, 7.233013851209895e-05, (1 * Dyne).Poundals())

	assertFloatEqual(t, 9.80665, (1 * KilogramForce).Newtons())
	assertFloatEqual(t, 980665, (1 * KilogramForce).Dynes())
	assertFloatEqual(t, 2.2046226218487757, (1 * KilogramForce).PoundForce())
	assertFloatEqual(t, 70.93163528396751, (1 * KilogramForce).Poundals())

	assertFloatEqual(t, 4.4482216152605, (1 * PoundForce).Newtons())
	assertFloatEqual(t, 444822.16152605, (1 * PoundForce).Dynes())
	assertFloatEqual(t, 0.45359237, (1 * PoundForce).KilogramForce())
	assertFloatEqual(t, 32.17404855643045, (1 * PoundForce).Poundals())

	assertFloatEqual(t, 0.138254954376, (1 * Poundal).Newtons())
	assertFloatEqual(t, 13825.4954376, (1 * Poundal).Dynes())
	assertFloatEqual(t, 0.014098081850173097, (1 * Poundal).KilogramForce())
	assertFloatEqual(t, 0.031080950171567253, (1 * Poundal).PoundForce())
}
