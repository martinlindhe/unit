package unit

import (
	"testing"
)

func TestEnergy(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptojoule).Yoctojoules())
	assertFloatEqual(t, 1e3, (1 * Attojoule).Zeptojoules())
	assertFloatEqual(t, 1e3, (1 * Femtojoule).Attojoules())
	assertFloatEqual(t, 1e3, (1 * Picojoule).Femtojoules())
	assertFloatEqual(t, 1e3, (1 * Nanojoule).Picojoules())
	assertFloatEqual(t, 1e3, (1 * Microjoule).Nanojoules())
	assertFloatEqual(t, 1e3, (1 * Millijoule).Microjoules())

	assertFloatEqual(t, 1e3, (1 * Joule).Millijoules())
	assertFloatEqual(t, 1e2, (1 * Joule).Centijoules())
	assertFloatEqual(t, 1e1, (1 * Joule).Decijoules())
	assertFloatEqual(t, 1e0, (1 * Joule).Joules())
	assertFloatEqual(t, 1e-1, (1 * Joule).Decajoules())
	assertFloatEqual(t, 1e-2, (1 * Joule).Hectojoules())
	assertFloatEqual(t, 1e-3, (1 * Joule).Kilojoules())

	assertFloatEqual(t, 1e-3, (1 * Kilojoule).Megajoules())
	assertFloatEqual(t, 1e-3, (1 * Megajoule).Gigajoules())
	assertFloatEqual(t, 1e-3, (1 * Gigajoule).Terajoules())
	assertFloatEqual(t, 1e-3, (1 * Terajoule).Petajoules())
	assertFloatEqual(t, 1e-3, (1 * Petajoule).Exajoules())
	assertFloatEqual(t, 1e-3, (1 * Exajoule).Zettajoules())
	assertFloatEqual(t, 1e-3, (1 * Zettajoule).Yottajoules())

	// watt hours
	assertFloatEqual(t, 1e3, (1 * ZeptowattHour).YoctowattHours())
	assertFloatEqual(t, 1e3, (1 * AttowattHour).ZeptowattHours())
	assertFloatEqual(t, 1e3, (1 * FemtowattHour).AttowattHours())
	assertFloatEqual(t, 1e3, (1 * PicowattHour).FemtowattHours())
	assertFloatEqual(t, 1e3, (1 * NanowattHour).PicowattHours())
	assertFloatEqual(t, 1e3, (1 * MicrowattHour).NanowattHours())
	assertFloatEqual(t, 1e3, (1 * MilliwattHour).MicrowattHours())

	assertFloatEqual(t, 1e3, (1 * WattHour).MilliwattHours())
	assertFloatEqual(t, 1e2, (1 * WattHour).CentiwattHours())
	assertFloatEqual(t, 1e1, (1 * WattHour).DeciwattHours())

	assertFloatEqual(t, 3600.0, (1 * WattHour).Joules())
	assertFloatEqual(t, 1e0, (1 * WattHour).WattHours())

	assertFloatEqual(t, 1e-1, (1 * WattHour).DecawattHours())
	assertFloatEqual(t, 1e-2, (1 * WattHour).HectowattHours())
	assertFloatEqual(t, 1e-3, (1 * WattHour).KilowattHours())

	assertFloatEqual(t, 1e-3, (1 * KilowattHour).MegawattHours())
	assertFloatEqual(t, 1e-3, (1 * MegawattHour).GigawattHours())
	assertFloatEqual(t, 1e-3, (1 * GigawattHour).TerawattHours())
	assertFloatEqual(t, 1e-3, (1 * TerawattHour).PetawattHours())
	assertFloatEqual(t, 1e-3, (1 * PetawattHour).ExawattHours())
	assertFloatEqual(t, 1e-3, (1 * ExawattHour).ZettawattHours())
	assertFloatEqual(t, 1e-3, (1 * ZettawattHour).YottawattHours())

	// calories
	assertFloatEqual(t, 4.184, (1 * Gramcalorie).Joules())
	assertFloatEqual(t, 4184.0, (1 * Kilocalorie).Joules())
	assertFloatEqual(t, 4.184e+6, (1 * Megacalorie).Joules())
}
