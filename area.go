package unit

// Area represents a SI unit of area (in square meters, m²)
type Area Unit

// ...
const (
	// SI
	SquareYoctometer      = SquareMeter * 1e-48
	SquareZeptometer      = SquareMeter * 1e-42
	SquareAttometer       = SquareMeter * 1e-36
	SquareFemtometer      = SquareMeter * 1e-30
	SquarePicometer       = SquareMeter * 1e-24
	SquareNanometer       = SquareMeter * 1e-18
	SquareMicrometer      = SquareMeter * 1e-12
	SquareMillimeter      = SquareMeter * 1e-6
	SquareCentimeter      = SquareMeter * 1e-4
	SquareDecimeter       = SquareMeter * 1e-2
	SquareMeter      Area = 1e0
	SquareDecameter       = SquareMeter * 1e2
	SquareHectometer      = SquareMeter * 1e4
	SquareKilometer       = SquareMeter * 1e6
	SquareMegameter       = SquareMeter * 1e12
	SquareGigameter       = SquareMeter * 1e18
	SquareTerameter       = SquareMeter * 1e24
	SquarePetameter       = SquareMeter * 1e30
	SquareExameter        = SquareMeter * 1e36
	SquareZettameter      = SquareMeter * 1e42
	SquareYottameter      = SquareMeter * 1e48

	// US
	SquareInch = SquareMeter * 0.00064516
	SquareFoot = SquareInch * 144
	SquareYard = SquareFoot * 9
	Acre       = SquareYard * 4840
	SquareMile = Acre * 640

	// TW
	//1坪(twPing)= 3.30579 平方公尺
	//1公頃(hectare,twGongQing)= 10000 平方公尺
	//1甲(twJia)= 2934*3.30579 平方公尺
	//1分(twFen)= 293.4*3.30579 平方公尺
	TWGongQing = SquareMeter * 10000
	TWPing     = SquareMeter / 3.30579
	TWFen      = SquareMeter / (293.4 * 3.30579)
	TWJia      = SquareMeter / (2934 * 3.30579)

	// imperial
	SquareRod = SquareFoot * 272.25
	Rood      = SquareYard * 1210

	// aliases
	Centiare    = SquareMeter
	Are         = SquareDecameter
	Hectare     = SquareHectometer
	SquarePerch = SquareRod
)

// SquareYoctometers returns the area in ym²
func (a Area) SquareYoctometers() float64 {
	return float64(a / SquareYoctometer)
}

// SquareZeptometers returns the area in zm²
func (a Area) SquareZeptometers() float64 {
	return float64(a / SquareZeptometer)
}

// SquareAttometers returns the area in am²
func (a Area) SquareAttometers() float64 {
	return float64(a / SquareAttometer)
}

// SquareFemtometers returns the area in fm²
func (a Area) SquareFemtometers() float64 {
	return float64(a / SquareFemtometer)
}

// SquarePicometers returns the area in pm²
func (a Area) SquarePicometers() float64 {
	return float64(a / SquarePicometer)
}

// SquareNanometers returns the area in nm²
func (a Area) SquareNanometers() float64 {
	return float64(a / SquareNanometer)
}

// SquareMicrometers returns the area in µm²
func (a Area) SquareMicrometers() float64 {
	return float64(a / SquareMicrometer)
}

// SquareMillimeters returns the area in mm²
func (a Area) SquareMillimeters() float64 {
	return float64(a / SquareMillimeter)
}

// SquareCentimeters returns the area in cm²
func (a Area) SquareCentimeters() float64 {
	return float64(a / SquareCentimeter)
}

// SquareDecimeters returns the area in dm²
func (a Area) SquareDecimeters() float64 {
	return float64(a / SquareDecimeter)
}

// SquareMeters returns the area in m²
func (a Area) SquareMeters() float64 {
	return float64(a / SquareMeter)
}

// Centiares returns the area in ca² (m²)
func (a Area) Centiares() float64 {
	return float64(a / Centiare)
}

// Ares returns the area in a²
func (a Area) Ares() float64 {
	return float64(a / Are)
}

// SquareDecameter returns the area in dam²
func (a Area) SquareDecameter() float64 {
	return float64(a / SquareDecameter)
}

// SquareHectometer returns the area in hm² (hectare)
func (a Area) SquareHectometer() float64 {
	return float64(a / SquareHectometer)
}

// Hectares returns the area in ha
func (a Area) Hectares() float64 {
	return float64(a / Hectare)
}

// SquareKilometers returns the area in km²
func (a Area) SquareKilometers() float64 {
	return float64(a / SquareKilometer)
}

// SquareMegameters returns the area in Mm²
func (a Area) SquareMegameters() float64 {
	return float64(a / SquareMegameter)
}

// SquareGigameters returns the area in Gm²
func (a Area) SquareGigameters() float64 {
	return float64(a / SquareGigameter)
}

// SquareTerameters returns the area in Tm²
func (a Area) SquareTerameters() float64 {
	return float64(a / SquareTerameter)
}

// SquarePetameters returns the area in Pm²
func (a Area) SquarePetameters() float64 {
	return float64(a / SquarePetameter)
}

// SquareExameters returns the area in Em²
func (a Area) SquareExameters() float64 {
	return float64(a / SquareExameter)
}

// SquareZettameters returns the area in Zm²
func (a Area) SquareZettameters() float64 {
	return float64(a / SquareZettameter)
}

// SquareYottameters returns the area in Ym²
func (a Area) SquareYottameters() float64 {
	return float64(a / SquareYottameter)
}

// SquareInches returns the area in in²
func (a Area) SquareInches() float64 {
	return float64(a / SquareInch)
}

// SquareFeet returns the area in ft²
func (a Area) SquareFeet() float64 {
	return float64(a / SquareFoot)
}

// SquareYards returns the area in yd
func (a Area) SquareYards() float64 {
	return float64(a / SquareYard)
}

// Acres returns the area in ac
func (a Area) Acres() float64 {
	return float64(a / Acre)
}

// SquareMiles returns the area in mi²
func (a Area) SquareMiles() float64 {
	return float64(a / SquareMile)
}

// SquareRods returns the area in square rods
func (a Area) SquareRods() float64 {
	return float64(a / SquareRod)
}

// Roods returns the area in roods
func (a Area) Roods() float64 {
	return float64(a / Rood)
}
