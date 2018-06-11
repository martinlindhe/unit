package unit

import (
	"testing"
)

func TestDatasize(t *testing.T) {

	// base 10 (SI prefixes)
	assertFloatEqual(t, 1e0, (1 * Bit).Bits())
	assertFloatEqual(t, 1e-3, (1 * Bit).Kilobits())
	assertFloatEqual(t, 1e-3, (1 * Kilobit).Megabits())
	assertFloatEqual(t, 1e-3, (1 * Megabit).Gigabits())
	assertFloatEqual(t, 1e-3, (1 * Gigabit).Terabits())
	assertFloatEqual(t, 1e-3, (1 * Terabit).Petabits())
	assertFloatEqual(t, 1e-3, (1 * Petabit).Exabits())
	assertFloatEqual(t, 1e-3, (1 * Exabit).Zettabits())
	assertFloatEqual(t, 1e-3, (1 * Zettabit).Yottabits())

	assertFloatEqual(t, 0.125, (1 * Bit).Bytes())
	assertFloatEqual(t, 1e0, (1 * Byte).Bytes())

	assertFloatEqual(t, 1e-3, (1 * Byte).Kilobytes())
	assertFloatEqual(t, 1e-3, (1 * Kilobyte).Megabytes())
	assertFloatEqual(t, 1e-3, (1 * Megabyte).Gigabytes())
	assertFloatEqual(t, 1e-3, (1 * Gigabyte).Terabytes())
	assertFloatEqual(t, 1e-3, (1 * Terabyte).Petabytes())
	assertFloatEqual(t, 1e-3, (1 * Petabyte).Exabytes())
	assertFloatEqual(t, 1e-3, (1 * Exabyte).Zettabytes())
	assertFloatEqual(t, 1e-3, (1 * Zettabyte).Yottabytes())

	// base 2 (IEC prefixes)
	assertFloatEqual(t, 0.0009765625, (1 * Bit).Kibibits())
	assertFloatEqual(t, 0.0009765625, (1 * Kibibit).Mebibits())
	assertFloatEqual(t, 0.0009765625, (1 * Mebibit).Gibibits())
	assertFloatEqual(t, 0.0009765625, (1 * Gibibit).Tebibits())
	assertFloatEqual(t, 0.0009765625, (1 * Tebibit).Pebibits())
	assertFloatEqual(t, 0.0009765625, (1 * Pebibit).Exbibits())
	assertFloatEqual(t, 0.0009765625, (1 * Exbibit).Zebibits())
	assertFloatEqual(t, 0.0009765625, (1 * Zebibit).Yobibits())

	assertFloatEqual(t, 0.0009765625, (1 * Byte).Kibibytes())
	assertFloatEqual(t, 0.0009765625, (1 * Kibibyte).Mebibytes())
	assertFloatEqual(t, 0.0009765625, (1 * Mebibyte).Gibibytes())
	assertFloatEqual(t, 0.0009765625, (1 * Gibibyte).Tebibytes())
	assertFloatEqual(t, 0.0009765625, (1 * Tebibyte).Pebibytes())
	assertFloatEqual(t, 0.0009765625, (1 * Pebibyte).Exbibytes())
	assertFloatEqual(t, 0.0009765625, (1 * Exbibyte).Zebibytes())
	assertFloatEqual(t, 0.0009765625, (1 * Zebibyte).Yobibytes())
}
