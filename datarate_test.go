package unit

import (
	"testing"
)

func TestDatarate(t *testing.T) {

	// base 10 (SI prefixes)
	assertFloatEqual(t, 1e0, (1 * BitPerSecond).BitsPerSecond())
	assertFloatEqual(t, 1e-3, (1 * BitPerSecond).KilobitsPerSecond())
	assertFloatEqual(t, 1e-3, (1 * KilobitPerSecond).MegabitsPerSecond())
	assertFloatEqual(t, 1e-3, (1 * MegabitPerSecond).GigabitsPerSecond())
	assertFloatEqual(t, 1e-3, (1 * GigabitPerSecond).TerabitsPerSecond())
	assertFloatEqual(t, 1e-3, (1 * TerabitPerSecond).PetabitsPerSecond())
	assertFloatEqual(t, 1e-3, (1 * PetabitPerSecond).ExabitsPerSecond())
	assertFloatEqual(t, 1e-3, (1 * ExabitPerSecond).ZettabitsPerSecond())
	assertFloatEqual(t, 1e-3, (1 * ZettabitPerSecond).YottabitsPerSecond())

	assertFloatEqual(t, 1e0, (1 * BytePerSecond).BytesPerSecond())
	assertFloatEqual(t, 1e-3, (1 * BytePerSecond).KilobytesPerSecond())
	assertFloatEqual(t, 1e-3, (1 * KilobytePerSecond).MegabytesPerSecond())
	assertFloatEqual(t, 1e-3, (1 * MegabytePerSecond).GigabytesPerSecond())
	assertFloatEqual(t, 1e-3, (1 * GigabytePerSecond).TerabytesPerSecond())
	assertFloatEqual(t, 1e-3, (1 * TerabytePerSecond).PetabytesPerSecond())
	assertFloatEqual(t, 1e-3, (1 * PetabytePerSecond).ExabytesPerSecond())
	assertFloatEqual(t, 1e-3, (1 * ExabytePerSecond).ZettabytesPerSecond())
	assertFloatEqual(t, 1e-3, (1 * ZettabytePerSecond).YottabytesPerSecond())

	// base 2 (IEC prefixes)
	assertFloatEqual(t, 0.0009765625, (1 * BitPerSecond).KibibitsPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * KibibitPerSecond).MebibitsPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * MebibitPerSecond).GibibitsPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * GibibitPerSecond).TebibitsPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * TebibitPerSecond).PebibitsPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * PebibitPerSecond).ExbibitsPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * ExbibitPerSecond).ZebibitsPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * ZebibitPerSecond).YobibitsPerSecond())

	assertFloatEqual(t, 0.0009765625, (1 * BytePerSecond).KibibytesPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * KibibytePerSecond).MebibytesPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * MebibytePerSecond).GibibytesPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * GibibytePerSecond).TebibytesPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * TebibytePerSecond).PebibytesPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * PebibytePerSecond).ExbibytesPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * ExbibytePerSecond).ZebibytesPerSecond())
	assertFloatEqual(t, 0.0009765625, (1 * ZebibytePerSecond).YobibytesPerSecond())
}
