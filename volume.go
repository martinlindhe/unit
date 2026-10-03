package unit

// Volume represents a volume in cubic meters
type Volume Unit

// ...
const (
	// SI
	CubicYoctometer        = CubicMeter * 1e-72
	CubicZeptometer        = CubicMeter * 1e-63
	CubicAttometer         = CubicMeter * 1e-54
	CubicFemtometer        = CubicMeter * 1e-45
	CubicPicometer         = CubicMeter * 1e-36
	CubicNanometer         = CubicMeter * 1e-27
	CubicMicrometer        = CubicMeter * 1e-18
	CubicMillimeter        = CubicMeter * 1e-9
	CubicCentimeter        = CubicMeter * 1e-6
	CubicDecimeter         = CubicMeter * 1e-3
	CubicMeter      Volume = 1e0
	CubicDecameter         = CubicMeter * 1e3
	CubicHectometer        = CubicMeter * 1e6
	CubicKilometer         = CubicMeter * 1e9
	CubicMegameter         = CubicMeter * 1e18
	CubicGigameter         = CubicMeter * 1e27
	CubicTerameter         = CubicMeter * 1e36
	CubicPetameter         = CubicMeter * 1e45
	CubicExameter          = CubicMeter * 1e54
	CubicZettameter        = CubicMeter * 1e63
	CubicYottameter        = CubicMeter * 1e72

	// SI derived
	Yoctoliter = Liter * 1e-24
	Zepoliter  = Liter * 1e-21
	Attoliter  = Liter * 1e-18
	Femtoliter = Liter * 1e-15
	Picoliter  = Liter * 1e-12
	Nanoliter  = Liter * 1e-9
	Microliter = Liter * 1e-6
	Milliliter = Liter * 1e-3
	Centiliter = Liter * 1e-2
	Deciliter  = Liter * 1e-1
	Liter      = CubicMeter * 1e-3
	Decaliter  = Liter * 1e1
	Hectoliter = Liter * 1e2
	Kiloliter  = Liter * 1e3
	Megaliter  = Liter * 1e6
	Gigaliter  = Liter * 1e9
	Teraliter  = Liter * 1e12
	Petaliter  = Liter * 1e15
	Exaliter   = Liter * 1e18
	Zettaliter = Liter * 1e21
	Yottaliter = Liter * 1e24

	// US
	CubicInch    = Milliliter * 16.387064
	CubicFoot    = CubicInch * 1728
	CubicYard    = CubicFoot * 27
	CubicMile    = CubicYard * 5451776000
	CubicFurlong = CubicMile * 0.00195314

	// imperial liquid
	ImperialGallon     = Liter * 4.54609
	ImperialQuart      = ImperialGallon / 4
	ImperialPint       = ImperialQuart / 2
	ImperialCup        = ImperialPint / 2
	ImperialGill       = ImperialPint / 4
	ImperialFluidOunce = ImperialGill / 5
	ImperialFluidDram  = ImperialFluidOunce / 8
	ImperialPeck       = ImperialGallon * 2
	ImperialBushel     = ImperialPeck * 4

	// metric cooking
	MetricTableSpoon = Milliliter * 15
	MetricTeaSpoon   = Milliliter * 5

	// US liquid
	USLiquidGallon = CubicInch * 231
	USLiquidQuart  = CubicInch * 57.75
	USLiquidPint   = CubicInch * 28.875
	USCup          = USLiquidPint / 2
	USLegalCup     = Milliliter * 240
	USGill         = Milliliter * 118.29411825
	USFluidDram    = USFluidOunce / 8
	USFluidOunce   = USLiquidGallon / 128
	USTableSpoon   = USFluidOunce / 2
	USTeaSpoon     = USTableSpoon / 3

	// US dry
	USDryQuart  = USDryGallon / 4
	USBushel    = USPeck * 4
	USPeck      = USDryGallon * 2
	USDryGallon = CubicInch * 268.8025
	USDryPint   = CubicInch * 33.6003125

	// misc
	AustralianTableSpoon = Milliliter * 20

	// aliases
	ImperialTableSpoon = MetricTableSpoon
	ImperialTeaSpoon   = MetricTeaSpoon
)

// Yoctoliters returns the volume in yl
func (v Volume) Yoctoliters() float64 {
	return float64(v / Yoctoliter)
}

// Zepoliters returns the volume in zl
func (v Volume) Zepoliters() float64 {
	return float64(v / Zepoliter)
}

// Attoliters returns the volume in al
func (v Volume) Attoliters() float64 {
	return float64(v / Attoliter)
}

// Femtoliters returns the volume in fl
func (v Volume) Femtoliters() float64 {
	return float64(v / Femtoliter)
}

// Picoliters returns the volume in pl
func (v Volume) Picoliters() float64 {
	return float64(v / Picoliter)
}

// Nanoliters returns the volume in nl
func (v Volume) Nanoliters() float64 {
	return float64(v / Nanoliter)
}

// Microliters returns the volume in µl
func (v Volume) Microliters() float64 {
	return float64(v / Microliter)
}

// Milliliters returns the volume in ml
func (v Volume) Milliliters() float64 {
	return float64(v / Milliliter)
}

// Centiliters returns the volume in cl
func (v Volume) Centiliters() float64 {
	return float64(v / Centiliter)
}

// Deciliters returns the volume in dl
func (v Volume) Deciliters() float64 {
	return float64(v / Deciliter)
}

// Liters returns the volume in l
func (v Volume) Liters() float64 {
	return float64(v / Liter)
}

// Decaliters returns the volume in Dl
func (v Volume) Decaliters() float64 {
	return float64(v / Decaliter)
}

// Hectoliters returns the volume in Hl
func (v Volume) Hectoliters() float64 {
	return float64(v / Hectoliter)
}

// Kiloliters returns the volume in Kl
func (v Volume) Kiloliters() float64 {
	return float64(v / Kiloliter)
}

// Megaliters returns the volume in Ml
func (v Volume) Megaliters() float64 {
	return float64(v / Megaliter)
}

// Gigaliters returns the volume in Gl
func (v Volume) Gigaliters() float64 {
	return float64(v / Gigaliter)
}

// Teraliters returns the volume in Tl
func (v Volume) Teraliters() float64 {
	return float64(v / Teraliter)
}

// Petaliters returns the volume in Pl
func (v Volume) Petaliters() float64 {
	return float64(v / Petaliter)
}

// Exaliters returns the volume in El
func (v Volume) Exaliters() float64 {
	return float64(v / Exaliter)
}

// Zettaliters returns the volume in Zl
func (v Volume) Zettaliters() float64 {
	return float64(v / Zettaliter)
}

// CubicYoctometers returns the volume in ym³
func (v Volume) CubicYoctometers() float64 {
	return float64(v / CubicYoctometer)
}

// CubicZeptometers returns the volume in zm³
func (v Volume) CubicZeptometers() float64 {
	return float64(v / CubicZeptometer)
}

// CubicAttometers returns the volume in am³
func (v Volume) CubicAttometers() float64 {
	return float64(v / CubicAttometer)
}

// CubicFemtometers returns the volume in fm³
func (v Volume) CubicFemtometers() float64 {
	return float64(v / CubicFemtometer)
}

// CubicPicometers returns the volume in pm³
func (v Volume) CubicPicometers() float64 {
	return float64(v / CubicPicometer)
}

// CubicNanometers returns the volume in nm³
func (v Volume) CubicNanometers() float64 {
	return float64(v / CubicNanometer)
}

// CubicMicrometers returns the volume in µm³
func (v Volume) CubicMicrometers() float64 {
	return float64(v / CubicMicrometer)
}

// CubicMillimeters returns the volume in mm³
func (v Volume) CubicMillimeters() float64 {
	return float64(v / CubicMillimeter)
}

// CubicCentimeters returns the volume in cm³
func (v Volume) CubicCentimeters() float64 {
	return float64(v / CubicCentimeter)
}

// CubicDecimeters returns the volume in dm³
func (v Volume) CubicDecimeters() float64 {
	return float64(v / CubicDecimeter)
}

// CubicMeters returns the volume in m³
func (v Volume) CubicMeters() float64 {
	return float64(v / CubicMeter)
}

// CubicDecameters returns the volume in dam³
func (v Volume) CubicDecameters() float64 {
	return float64(v / CubicDecameter)
}

// CubicHectometers returns the volume in hm³
func (v Volume) CubicHectometers() float64 {
	return float64(v / CubicHectometer)
}

// CubicKilometers returns the volume in km³
func (v Volume) CubicKilometers() float64 {
	return float64(v / CubicKilometer)
}

// CubicMegameters returns the volume in Mm³
func (v Volume) CubicMegameters() float64 {
	return float64(v / CubicMegameter)
}

// CubicGigameters returns the volume in Gm³
func (v Volume) CubicGigameters() float64 {
	return float64(v / CubicGigameter)
}

// CubicTerameters returns the volume in Tm³
func (v Volume) CubicTerameters() float64 {
	return float64(v / CubicTerameter)
}

// CubicPetameters returns the volume in Pm³
func (v Volume) CubicPetameters() float64 {
	return float64(v / CubicPetameter)
}

// CubicExameters returns the volume in Em³
func (v Volume) CubicExameters() float64 {
	return float64(v / CubicExameter)
}

// CubicZettameters returns the volume in Zm³
func (v Volume) CubicZettameters() float64 {
	return float64(v / CubicZettameter)
}

// CubicYottameters returns the volume in Ym³
func (v Volume) CubicYottameters() float64 {
	return float64(v / CubicYottameter)
}

// CubicInches returns the volume in in³
func (v Volume) CubicInches() float64 {
	return float64(v / CubicInch)
}

// CubicFeet returns the volume in ft³
func (v Volume) CubicFeet() float64 {
	return float64(v / CubicFoot)
}

// CubicYards returns the volume in yd³
func (v Volume) CubicYards() float64 {
	return float64(v / CubicYard)
}

// CubicMiles returns the volume in mi³
func (v Volume) CubicMiles() float64 {
	return float64(v / CubicMile)
}

// CubicFurlongs returns the volume in furlong³
func (v Volume) CubicFurlongs() float64 {
	return float64(v / CubicFurlong)
}

// ImperialGallons returns the volume in imperial gallons
func (v Volume) ImperialGallons() float64 {
	return float64(v / ImperialGallon)
}

// ImperialQuarts returns the volume in imperial quarts
func (v Volume) ImperialQuarts() float64 {
	return float64(v / ImperialQuart)
}

// ImperialPints returns the volume in imperial pints
func (v Volume) ImperialPints() float64 {
	return float64(v / ImperialPint)
}

// ImperialGills returns the volume in imperial gills
func (v Volume) ImperialGills() float64 {
	return float64(v / ImperialGill)
}

// ImperialCups returns the volume in imperial cups
func (v Volume) ImperialCups() float64 {
	return float64(v / ImperialCup)
}

// ImperialFluidOunces returns the volume in imperial fluid ounces
func (v Volume) ImperialFluidOunces() float64 {
	return float64(v / ImperialFluidOunce)
}

// ImperialFluidDrams returns the volume in imperial fluid drams
func (v Volume) ImperialFluidDrams() float64 {
	return float64(v / ImperialFluidDram)
}

// ImperialPecks returns the volume in imperial pecks
func (v Volume) ImperialPecks() float64 {
	return float64(v / ImperialPeck)
}

// ImperialBushels returns the volume in imperial bushels
func (v Volume) ImperialBushels() float64 {
	return float64(v / ImperialBushel)
}

// MetricTableSpoons returns the volume in metric/imperial tablespoons
func (v Volume) MetricTableSpoons() float64 {
	return float64(v / MetricTableSpoon)
}

// MetricTeaSpoons returns the volume in metric/imperial teaspoons
func (v Volume) MetricTeaSpoons() float64 {
	return float64(v / MetricTeaSpoon)
}

// ImperialTableSpoons returns the volume in metric/imperial tablespoons
func (v Volume) ImperialTableSpoons() float64 {
	return v.MetricTableSpoons()
}

// ImperialTeaSpoons returns the volume in metric/imperial teaspoons
func (v Volume) ImperialTeaSpoons() float64 {
	return v.MetricTeaSpoons()
}

// AustralianTableSpoons returns the volume in Australian tablespoons
func (v Volume) AustralianTableSpoons() float64 {
	return float64(v / AustralianTableSpoon)
}

// USLiquidGallons returns the volume in US liquid gallons
func (v Volume) USLiquidGallons() float64 {
	return float64(v / USLiquidGallon)
}

// USLiquidQuarts returns the volume in US liquid quarts
func (v Volume) USLiquidQuarts() float64 {
	return float64(v / USLiquidQuart)
}

// USLiquidPints returns the volume in US liquid pints
func (v Volume) USLiquidPints() float64 {
	return float64(v / USLiquidPint)
}

// USCups returns the volume in US cups
func (v Volume) USCups() float64 {
	return float64(v / USCup)
}

// USLegalCups returns the volume in US legal cups
func (v Volume) USLegalCups() float64 {
	return float64(v / USLegalCup)
}

// USGills returns the volume in US gills
func (v Volume) USGills() float64 {
	return float64(v / USGill)
}

// USTableSpoons returns the volume in US table spoons
func (v Volume) USTableSpoons() float64 {
	return float64(v / USTableSpoon)
}

// USTeaSpoons returns the volume in US tea spoons
func (v Volume) USTeaSpoons() float64 {
	return float64(v / USTeaSpoon)
}

// USFluidDrams returns the volume in US fluid drams
func (v Volume) USFluidDrams() float64 {
	return float64(v / USFluidDram)
}

// USFluidOunces returns the volume in US fluid ounces
func (v Volume) USFluidOunces() float64 {
	return float64(v / USFluidOunce)
}

// USDryQuarts returns the volume in US dry quarts
func (v Volume) USDryQuarts() float64 {
	return float64(v / USDryQuart)
}

// USBushels returns the volume in US bushels
func (v Volume) USBushels() float64 {
	return float64(v / USBushel)
}

// USPecks returns the volume in US pecks
func (v Volume) USPecks() float64 {
	return float64(v / USPeck)
}

// USDryGallons returns the volume in US dry gallons
func (v Volume) USDryGallons() float64 {
	return float64(v / USDryGallon)
}

// USDryPints returns the volume in US dry pints
func (v Volume) USDryPints() float64 {
	return float64(v / USDryPint)
}
