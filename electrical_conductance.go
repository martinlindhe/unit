package unit

// ElectricalConductance represents a SI unit of electrical conductance (in siemens, S)
type ElectricalConductance Unit

// ...
const (
	Siemens ElectricalConductance = 1e0 // SI
)

// Siemens returns the amount of electrical conductance in S
func (e ElectricalConductance) Siemens() float64 {
	return float64(e)
}
