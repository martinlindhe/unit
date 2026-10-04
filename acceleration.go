package unit

// Acceleration represents a SI unit of acceleration (in meter per second squared, m/s²)
type Acceleration Unit

// ...
const (
	CentimeterPerSecondSquared              = MeterPerSecondSquared * 1e-2         // SI
	MeterPerSecondSquared      Acceleration = 1e0                                  // SI
	FootPerSecondSquared                    = MeterPerSecondSquared * 3048 / 10000 // US
	StandardGravity                         = MeterPerSecondSquared * 9.80665      // space
	Gal                                     = CentimeterPerSecondSquared           // alias
)

// CentimetersPerSecondSquared returns the acceleration in cm/s²
func (a Acceleration) CentimetersPerSecondSquared() float64 {
	return float64(a / CentimeterPerSecondSquared)
}

// MetersPerSecondSquared returns the acceleration in m/s²
func (a Acceleration) MetersPerSecondSquared() float64 {
	return float64(a / MeterPerSecondSquared)
}

// FeetPerSecondSquared returns the acceleration in ft/s²
func (a Acceleration) FeetPerSecondSquared() float64 {
	return float64(a / FootPerSecondSquared)
}

// StandardGravity returns the acceleration in ɡ₀
func (a Acceleration) StandardGravity() float64 {
	return float64(a / StandardGravity)
}

// Gals returns the acceleration in Gal
func (a Acceleration) Gals() float64 {
	return float64(a / Gal)
}
