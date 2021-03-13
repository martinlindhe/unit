package unit

// Pressure represents a SI derived unit of pressure (in pascal, Pa)
type Pressure Unit

// ...
const (
	// SI derived
	Yoctopascal          = Pascal * 1e-24
	Zeptopascal          = Pascal * 1e-21
	Attopascal           = Pascal * 1e-18
	Femtopascal          = Pascal * 1e-15
	Picopascal           = Pascal * 1e-12
	Nanopascal           = Pascal * 1e-9
	Micropascal          = Pascal * 1e-6
	Millipascal          = Pascal * 1e-3
	Centipascal          = Pascal * 1e-2
	Decipascal           = Pascal * 1e-1
	Pascal      Pressure = 1e0
	Decapascal           = Pascal * 1e1
	Hectopascal          = Pascal * 1e2
	Kilopascal           = Pascal * 1e3
	Megapascal           = Pascal * 1e6
	Gigapascal           = Pascal * 1e9
	Terapascal           = Pascal * 1e12
	Petapascal           = Pascal * 1e15
	Exapascal            = Pascal * 1e18
	Zettapascal          = Pascal * 1e21
	Yottapascal          = Pascal * 1e24

	// non-SI
	Yoctobar = Bar * 1e-24
	Zeptobar = Bar * 1e-21
	Attobar  = Bar * 1e-18
	Femtobar = Bar * 1e-15
	Picobar  = Bar * 1e-12
	Nanobar  = Bar * 1e-9
	Microbar = Bar * 1e-6
	Millibar = Bar * 1e-3
	Centibar = Bar * 1e-2
	Decibar  = Bar * 1e-1
	Bar      = Pascal * 1e5
	Decabar  = Bar * 1e1
	Hectobar = Bar * 1e2
	Kilobar  = Bar * 1e3
	Megabar  = Bar * 1e6
	Gigabar  = Bar * 1e9
	Terabar  = Bar * 1e12
	Petabar  = Bar * 1e15
	Exabar   = Bar * 1e18
	Zettabar = Bar * 1e21
	Yottabar = Bar * 1e24

	Atmosphere          = Pascal * 1.01325 * 1e5
	TechAtmosphere      = Pascal * 9.80665 * 1e4
	Torr                = Pascal * 133.3224
	PoundsPerSquareInch = Pascal * 6.8948 * 1e3
	InchOfMercury       = Pascal * 3386.389
)

// Yoctopascals returns the pressure in yPa
func (p Pressure) Yoctopascals() float64 {
	return float64(p / Yoctopascal)
}

// Zeptopascals returns the pressure in zPa
func (p Pressure) Zeptopascals() float64 {
	return float64(p / Zeptopascal)
}

// Attopascals returns the pressure in aPa
func (p Pressure) Attopascals() float64 {
	return float64(p / Attopascal)
}

// Femtopascals returns the pressure in fPa
func (p Pressure) Femtopascals() float64 {
	return float64(p / Femtopascal)
}

// Picopascals returns the pressure in pPa
func (p Pressure) Picopascals() float64 {
	return float64(p / Picopascal)
}

// Nanopascals returns the pressure in nPa
func (p Pressure) Nanopascals() float64 {
	return float64(p / Nanopascal)
}

// Micropascals returns the pressure in µPa
func (p Pressure) Micropascals() float64 {
	return float64(p / Micropascal)
}

// Millipascals returns the pressure in mPa
func (p Pressure) Millipascals() float64 {
	return float64(p / Millipascal)
}

// Centipascals returns the pressure in cPa
func (p Pressure) Centipascals() float64 {
	return float64(p / Centipascal)
}

// Decipascals returns the pressure in dPa
func (p Pressure) Decipascals() float64 {
	return float64(p / Decipascal)
}

// Pascals returns the pressure in Pa
func (p Pressure) Pascals() float64 {
	return float64(p)
}

// Decapascals returns the pressure in daPa
func (p Pressure) Decapascals() float64 {
	return float64(p / Decapascal)
}

// Hectopascals returns the pressure in hPa
func (p Pressure) Hectopascals() float64 {
	return float64(p / Hectopascal)
}

// Kilopascals returns the pressure in kPa
func (p Pressure) Kilopascals() float64 {
	return float64(p / Kilopascal)
}

// Megapascals returns the pressure in MPa
func (p Pressure) Megapascals() float64 {
	return float64(p / Megapascal)
}

// Gigapascals returns the pressure in GPa
func (p Pressure) Gigapascals() float64 {
	return float64(p / Gigapascal)
}

// Terapascals returns the pressure in TPa
func (p Pressure) Terapascals() float64 {
	return float64(p / Terapascal)
}

// Petapascals returns the pressure in PPa
func (p Pressure) Petapascals() float64 {
	return float64(p / Petapascal)
}

// Exapascals returns the pressure in EPa
func (p Pressure) Exapascals() float64 {
	return float64(p / Exapascal)
}

// Zettapascals returns the pressure in ZPa
func (p Pressure) Zettapascals() float64 {
	return float64(p / Zettapascal)
}

// Yottapascals returns the pressure in YPa
func (p Pressure) Yottapascals() float64 {
	return float64(p / Yottapascal)
}

// Yoctobars returns the pressure in ybar
func (p Pressure) Yoctobars() float64 {
	return float64(p / Yoctobar)
}

// Zeptobars returns the pressure in zbar
func (p Pressure) Zeptobars() float64 {
	return float64(p / Zeptobar)
}

// Attobars returns the pressure in abar
func (p Pressure) Attobars() float64 {
	return float64(p / Attobar)
}

// Femtobars returns the pressure in fbar
func (p Pressure) Femtobars() float64 {
	return float64(p / Femtobar)
}

// Picobars returns the pressure in pbar
func (p Pressure) Picobars() float64 {
	return float64(p / Picobar)
}

// Nanobars returns the pressure in nbar
func (p Pressure) Nanobars() float64 {
	return float64(p / Nanobar)
}

// Microbars returns the pressure in µbar
func (p Pressure) Microbars() float64 {
	return float64(p / Microbar)
}

// Millibars returns the pressure in mbar
func (p Pressure) Millibars() float64 {
	return float64(p / Millibar)
}

// Centibars returns the pressure in cbar
func (p Pressure) Centibars() float64 {
	return float64(p / Centibar)
}

// Decibars returns the pressure in dbar
func (p Pressure) Decibars() float64 {
	return float64(p / Decibar)
}

// Bars returns the pressure in bar
func (p Pressure) Bars() float64 {
	return float64(p / Bar)
}

// Decabars returns the pressure in dabar
func (p Pressure) Decabars() float64 {
	return float64(p / Decabar)
}

// Hectobars returns the pressure in hbar
func (p Pressure) Hectobars() float64 {
	return float64(p / Hectobar)
}

// Kilobars returns the pressure in kbar
func (p Pressure) Kilobars() float64 {
	return float64(p / Kilobar)
}

// Megabars returns the pressure in Mbar
func (p Pressure) Megabars() float64 {
	return float64(p / Megabar)
}

// Gigabars returns the pressure in Gbar
func (p Pressure) Gigabars() float64 {
	return float64(p / Gigabar)
}

// Terabars returns the pressure in Tbar
func (p Pressure) Terabars() float64 {
	return float64(p / Terabar)
}

// Petabars returns the pressure in Pbar
func (p Pressure) Petabars() float64 {
	return float64(p / Petabar)
}

// Exabars returns the pressure in Ebar
func (p Pressure) Exabars() float64 {
	return float64(p / Exabar)
}

// Zettabars returns the pressure in Zbar
func (p Pressure) Zettabars() float64 {
	return float64(p / Zettabar)
}

// Yottabars returns the pressure in Ybar
func (p Pressure) Yottabars() float64 {
	return float64(p / Yottabar)
}

// Atmospheres returns the pressure in atm
func (p Pressure) Atmospheres() float64 {
	return float64(p / Atmosphere)
}

// TechAtmospheres returns the pressure in at
func (p Pressure) TechAtmospheres() float64 {
	return float64(p / TechAtmosphere)
}

// Torrs returns the pressure in Torr
func (p Pressure) Torrs() float64 {
	return float64(p / Torr)
}

// PoundsPerSquareInch returns the pressure in psi
func (p Pressure) PoundsPerSquareInch() float64 {
	return float64(p / PoundsPerSquareInch)
}

// InchOfMercury returns the pressure in inch of mercury
func (p Pressure) InchOfMercury() float64 {
	return float64(p / InchOfMercury)
}
