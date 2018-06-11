package unit

import (
	"testing"
)

func TestSpeed(t *testing.T) {

	// SI
	assertFloatEqual(t, 10.0, (10 * MetersPerSecond).MetersPerSecond())
	assertFloatEqual(t, 35.99997120002304, (10 * MetersPerSecond).KilometersPerHour())

	// US
	assertFloatEqual(t, 32.808398950131235, (10 * MetersPerSecond).FeetPerSecond())
	assertFloatEqual(t, 22.369362920544024, (10 * MetersPerSecond).MilesPerHour())

	// misc
	assertFloatEqual(t, 19.438461717893492, (10 * MetersPerSecond).Knots())

	// space
	assertFloatEqual(t, 0.03335640951981521, (10000000 * MetersPerSecond).SpeedOfLight())
}
