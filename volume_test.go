package unit

import (
	"testing"
)

func TestVolume(t *testing.T) {

	// SI derived
	assertFloatEqual(t, 1e3, (1 * Zepoliter).Yoctoliters())
	assertFloatEqual(t, 1e3, (1 * Attoliter).Zepoliters())
	assertFloatEqual(t, 1e3, (1 * Femtoliter).Attoliters())
	assertFloatEqual(t, 1e3, (1 * Picoliter).Femtoliters())
	assertFloatEqual(t, 1e3, (1 * Nanoliter).Picoliters())
	assertFloatEqual(t, 1e3, (1 * Microliter).Nanoliters())
	assertFloatEqual(t, 1e3, (1 * Milliliter).Microliters())

	assertFloatEqual(t, 1e3, (1 * Liter).Milliliters())
	assertFloatEqual(t, 1e2, (1 * Liter).Centiliters())
	assertFloatEqual(t, 1e1, (1 * Liter).Deciliters())
	assertFloatEqual(t, 1e0, (1 * Liter).Liters())
	assertFloatEqual(t, 1e-1, (1 * Liter).Decaliters())
	assertFloatEqual(t, 1e-2, (1 * Liter).Hectoliters())
	assertFloatEqual(t, 1e-3, (1 * Liter).Kiloliters())

	assertFloatEqual(t, 1e-3, (1 * Kiloliter).Megaliters())
	assertFloatEqual(t, 1e-3, (1 * Megaliter).Gigaliters())
	assertFloatEqual(t, 1e-3, (1 * Gigaliter).Teraliters())
	assertFloatEqual(t, 1e-3, (1 * Teraliter).Petaliters())
	assertFloatEqual(t, 1e-3, (1 * Petaliter).Exaliters())
	assertFloatEqual(t, 1e-3, (1 * Exaliter).Zettaliters())

	// SI
	assertFloatEqual(t, 1e9, (1 * CubicZeptometer).CubicYoctometers())
	assertFloatEqual(t, 1e9, (1 * CubicAttometer).CubicZeptometers())
	assertFloatEqual(t, 1e9, (1 * CubicFemtometer).CubicAttometers())
	assertFloatEqual(t, 1e9, (1 * CubicPicometer).CubicFemtometers())
	assertFloatEqual(t, 1e9, (1 * CubicNanometer).CubicPicometers())
	assertFloatEqual(t, 1e9, (1 * CubicMicrometer).CubicNanometers())
	assertFloatEqual(t, 1e9, (1 * CubicMillimeter).CubicMicrometers())

	assertFloatEqual(t, 1e9, (1 * CubicMeter).CubicMillimeters())
	assertFloatEqual(t, 1e6, (1 * CubicMeter).CubicCentimeters())
	assertFloatEqual(t, 1e3, (1 * CubicMeter).CubicDecimeters())

	assertFloatEqual(t, 1e-3, (1 * Liter).CubicMeters())
	assertFloatEqual(t, 1e0, (1 * CubicMeter).CubicMeters())

	assertFloatEqual(t, 1e-3, (1 * CubicMeter).CubicDecameters())
	assertFloatEqual(t, 1e-6, (1 * CubicMeter).CubicHectometers())
	assertFloatEqual(t, 1e-9, (1 * CubicMeter).CubicKilometers())

	assertFloatEqual(t, 1e-9, (1 * CubicKilometer).CubicMegameters())
	assertFloatEqual(t, 1e-9, (1 * CubicMegameter).CubicGigameters())
	assertFloatEqual(t, 1e-9, (1 * CubicGigameter).CubicTerameters())
	assertFloatEqual(t, 1e-9, (1 * CubicTerameter).CubicPetameters())
	assertFloatEqual(t, 1e-9, (1 * CubicPetameter).CubicExameters())
	assertFloatEqual(t, 1e-9, (1 * CubicExameter).CubicZettameters())
	assertFloatEqual(t, 1e-9, (1 * CubicZettameter).CubicYottameters())

	// US
	assertFloatEqual(t, 1728.0, (1 * CubicFoot).CubicInches())
	assertFloatEqual(t, 1e0, (1 * CubicFoot).CubicFeet())
	assertFloatEqual(t, 0.037037037037037035, (1 * CubicFoot).CubicYards())
	assertFloatEqual(t, 1.8342646506386176e-10, (1 * CubicYard).CubicMiles())
	assertFloatEqual(t, 9.391362885602761e-8, (1 * CubicYard).CubicFurlongs())

	// imperial liquid
	assertFloatEqual(t, 0.21996924829908776, (1 * Liter).ImperialGallons())
	assertFloatEqual(t, 0.879876993196351, (1 * Liter).ImperialQuarts())
	assertFloatEqual(t, 1.759753986392702, (1 * Liter).ImperialPints())
	assertFloatEqual(t, 7.039015945570808, (1 * Liter).ImperialGills())
	assertFloatEqual(t, 3.519507972785404, (1 * Liter).ImperialCups())
	assertFloatEqual(t, 35.19507972785404, (1 * Liter).ImperialFluidOunces())
	assertFloatEqual(t, 281.56063782283235, (1 * Liter).ImperialFluidDrams())
	assertFloatEqual(t, 0.10998462414954388, (1 * Liter).ImperialPecks())
	assertFloatEqual(t, 0.02749615603738597, (1 * Liter).ImperialBushels())

	// US liquid
	assertFloatEqual(t, 0.26417205235814845, (1 * Liter).USLiquidGallons())
	assertFloatEqual(t, 1.0566882094325938, (1 * Liter).USLiquidQuarts())
	assertFloatEqual(t, 2.1133764188651876, (1 * Liter).USLiquidPints())
	assertFloatEqual(t, 4.226752837730375, (1 * Liter).USCups())
	assertFloatEqual(t, 4.166666666666667, (1 * Liter).USLegalCups())
	assertFloatEqual(t, 8.45350567546075, (1 * Liter).USGills())
	assertFloatEqual(t, 67.628045403686, (1 * Liter).USTableSpoons())
	assertFloatEqual(t, 202.884136211058, (1 * Liter).USTeaSpoons())
	assertFloatEqual(t, 270.512181614744, (1 * Liter).USFluidDrams())
	assertFloatEqual(t, 33.814022701843, (1 * Liter).USFluidOunces())

	// US dry
	assertFloatEqual(t, 0.9080829842688559, (1 * Liter).USDryQuarts())
	assertFloatEqual(t, 0.028377593258401747, (1 * Liter).USBushels())
	assertFloatEqual(t, 0.11351037303360699, (1 * Liter).USPecks())
	assertFloatEqual(t, 0.22702074606721398, (1 * Liter).USDryGallons())
	assertFloatEqual(t, 1.8161659685377118, (1 * Liter).USDryPints())
}
