package unit

import (
	"testing"
)

func TestDuration(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptosecond).Yoctoseconds())
	assertFloatEqual(t, 1e3, (1 * Attosecond).Zeptoseconds())
	assertFloatEqual(t, 1e3, (1 * Femtosecond).Attoseconds())
	assertFloatEqual(t, 1e3, (1 * Picosecond).Femtoseconds())
	assertFloatEqual(t, 1e3, (1 * Nanosecond).Picoseconds())
	assertFloatEqual(t, 1e3, (1 * Microsecond).Nanoseconds())
	assertFloatEqual(t, 1e3, (1 * Millisecond).Microseconds())

	assertFloatEqual(t, 1e3, (1 * Second).Milliseconds())
	assertFloatEqual(t, 1e2, (1 * Second).Centiseconds())
	assertFloatEqual(t, 1e1, (1 * Second).Deciseconds())
	assertFloatEqual(t, 1e0, (1 * Second).Seconds())
	assertFloatEqual(t, 1e-1, (1 * Second).Decaseconds())
	assertFloatEqual(t, 1e-2, (1 * Second).Hectoseconds())
	assertFloatEqual(t, 1e-3, (1 * Second).Kiloseconds())

	assertFloatEqual(t, 1e-3, (1 * Kilosecond).Megaseconds())
	assertFloatEqual(t, 1e-3, (1 * Megasecond).Gigaseconds())
	assertFloatEqual(t, 1e-3, (1 * Gigasecond).Teraseconds())
	assertFloatEqual(t, 1e-3, (1 * Terasecond).Petaseconds())
	assertFloatEqual(t, 1e-3, (1 * Petasecond).Exaseconds())
	assertFloatEqual(t, 1e-3, (1 * Exasecond).Zettaseconds())
	assertFloatEqual(t, 1e-3, (1 * Zettasecond).Yottaseconds())

	// non-SI
	assertFloatEqual(t, 0.016666666666666666, (1 * Second).Minutes())
	assertFloatEqual(t, 0.0002777777777777778, (1 * Second).Hours())
	assertFloatEqual(t, 1.1574074074074073e-5, (1 * Second).Days())
	assertFloatEqual(t, 1.6534391534391535e-6, (1 * Second).Weeks())
	assertFloatEqual(t, 3.8580246913580245e-7, (1 * Second).ThirtyDayMonths())
	assertFloatEqual(t, 3.168808781402895e-8, (1 * Second).JulianYears())
}
