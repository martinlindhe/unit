package unit

import (
	"testing"
)

func TestForce(t *testing.T) {

	// constants from https://en.wikipedia.org/wiki/Force#Units_of_measurement
	assertFloatEqual(t, 1e5, (1 * Newton).Dynes())
	assertFloatEqual(t, 0.10197162129779283, (1 * Newton).KilogramForce())
	assertFloatEqual(t, 0.22480892365533914, (1 * Newton).PoundForce())
	assertFloatEqual(t, 7.233011464323171, (1 * Newton).Poundals())

	assertFloatEqual(t, 1e-5, (1 * Dyne).Newtons())
	assertFloatEqual(t, 1.0197162129779284e-6, (1 * Dyne).KilogramForce())
	assertFloatEqual(t, 2.2480892365533917e-6, (1 * Dyne).PoundForce())
	assertFloatEqual(t, 7.233011464323172e-05, (1 * Dyne).Poundals())

	assertFloatEqual(t, 9.80665, (1 * KilogramForce).Newtons())
	assertFloatEqual(t, 980665, (1 * KilogramForce).Dynes())
	assertFloatEqual(t, 2.204622431164631, (1 * KilogramForce).PoundForce())
	assertFloatEqual(t, 70.93161187660482, (1 * KilogramForce).Poundals())

	assertFloatEqual(t, 4.448222, (1 * PoundForce).Newtons())
	assertFloatEqual(t, 444822.2, (1 * PoundForce).Dynes())
	assertFloatEqual(t, 0.45359240923251065, (1 * PoundForce).KilogramForce())
	assertFloatEqual(t, 32.174040721854546, (1 * PoundForce).Poundals())

	assertFloatEqual(t, 0.138255, (1 * Poundal).Newtons())
	assertFloatEqual(t, 13825.5, (1 * Poundal).Dynes())
	assertFloatEqual(t, 0.014098086502526346, (1 * Poundal).KilogramForce())
	assertFloatEqual(t, 0.03108095773996891, (1 * Poundal).PoundForce())
}
