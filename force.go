package unit

// Force represents a SI unit of force (in newtons, N)
type Force Unit

// ...
const (
	// SI
	Newton Force = 1e0

	// non-SI
	Dyne          = Newton * 1e-5
	KilogramForce = Newton * Force(StandardGravity)
	PoundForce    = Newton * 4.448222
	Poundal       = Newton * 0.138255

	// aliases
	Kilopond = KilogramForce
)

// Newtons returns the force in N
func (f Force) Newtons() float64 {
	return float64(f)
}

// Dynes returns the force in dyn
func (f Force) Dynes() float64 {
	return float64(f / Dyne)
}

// KilogramForce returns the force in kp
func (f Force) KilogramForce() float64 {
	return float64(f / KilogramForce)
}

// PoundForce returns the force in lbf
func (f Force) PoundForce() float64 {
	return float64(f / PoundForce)
}

// Poundals returns the force in pdl
func (f Force) Poundals() float64 {
	return float64(f / Poundal)
}
