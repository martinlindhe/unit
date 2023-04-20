package unit

// Mass represents a SI unit of mass (in kilograms, kg)
type Mass Unit

// ...
const (
	// SI
	Yoctogram       = Gram * 1e-24
	Zeptogram       = Gram * 1e-21
	Attogram        = Gram * 1e-18
	Femtogram       = Gram * 1e-15
	Picogram        = Gram * 1e-12
	Nanogram        = Gram * 1e-9
	Microgram       = Gram * 1e-6
	Milligram       = Gram * 1e-3
	Centigram       = Gram * 1e-2
	Decigram        = Gram * 1e-1
	Gram            = Kilogram * 1e-3
	Decagram        = Gram * 1e1
	Hectogram       = Gram * 1e2
	Kilogram   Mass = 1e0
	Megagram        = Gram * 1e6
	Gigagram        = Gram * 1e9
	Teragram        = Gram * 1e12
	Petagram        = Gram * 1e15
	Exagram         = Gram * 1e18
	Zettagram       = Gram * 1e21
	Yottagram       = Gram * 1e24
	Ronnagram       = Gram * 1e27
	Quettagram      = Gram * 1e30

	// non-SI
	Tonne     = Megagram
	Kilotonne = Gigagram
	Megatonne = Teragram
	Gigatonne = Petagram
	Teratonne = Exagram
	Petatonne = Zettagram
	Exatonne  = Yottagram

	// US, avoirdupois
	TroyGrain          = Milligram * 64.79891
	AvoirdupoisDram    = AvoirdupoisOunce / 16
	AvoirdupoisOunce   = TroyGrain * 437.5
	AvoirdupoisPound   = TroyGrain * 7000
	UsStone            = AvoirdupoisPound * 14
	UsQuarter          = ShortHundredweight / 4
	ShortHundredweight = AvoirdupoisPound * 100

	// UK
	UkStone           = Gram * 6350.29318
	UkQuarter         = LongHundredweight / 4
	LongHundredweight = UkStone * 8
	TroyOunce         = TroyGrain * 480
	TroyPound         = TroyGrain * 5760

	// aliases
	CentalHundredweight   = ShortHundredweight // british
	ImperialHundredweight = LongHundredweight  // british
)

// Yoctograms returns the mass in yg
func (m Mass) Yoctograms() float64 {
	return float64(m / Yoctogram)
}

// Zeptograms returns the mass in zg
func (m Mass) Zeptograms() float64 {
	return float64(m / Zeptogram)
}

// Attograms returns the mass in ag
func (m Mass) Attograms() float64 {
	return float64(m / Attogram)
}

// Femtograms returns the mass in fg
func (m Mass) Femtograms() float64 {
	return float64(m / Femtogram)
}

// Picograms returns the mass in pg
func (m Mass) Picograms() float64 {
	return float64(m / Picogram)
}

// Nanograms returns the mass in ng
func (m Mass) Nanograms() float64 {
	return float64(m / Nanogram)
}

// Micrograms returns the mass in µg
func (m Mass) Micrograms() float64 {
	return float64(m / Microgram)
}

// Milligrams returns the mass in mg
func (m Mass) Milligrams() float64 {
	return float64(m / Milligram)
}

// Centigrams returns the mass in cg
func (m Mass) Centigrams() float64 {
	return float64(m / Centigram)
}

// Decigrams returns the mass in dg
func (m Mass) Decigrams() float64 {
	return float64(m / Decigram)
}

// Grams returns the mass in g
func (m Mass) Grams() float64 {
	return float64(m / Gram)
}

// Decagrams returns the mass in dag
func (m Mass) Decagrams() float64 {
	return float64(m / Decagram)
}

// Hectograms returns the mass in hg
func (m Mass) Hectograms() float64 {
	return float64(m / Hectogram)
}

// Kilograms returns the mass in kg
func (m Mass) Kilograms() float64 {
	return float64(m)
}

// Megagrams returns the mass in Mg
func (m Mass) Megagrams() float64 {
	return float64(m / Megagram)
}

// Gigagrams returns the mass in Gg
func (m Mass) Gigagrams() float64 {
	return float64(m / Gigagram)
}

// Teragrams returns the mass in Tg
func (m Mass) Teragrams() float64 {
	return float64(m / Teragram)
}

// Petagrams returns the mass in Pg
func (m Mass) Petagrams() float64 {
	return float64(m / Petagram)
}

// Exagrams returns the mass in Eg
func (m Mass) Exagrams() float64 {
	return float64(m / Exagram)
}

// Zettagrams returns the mass in Zg
func (m Mass) Zettagrams() float64 {
	return float64(m / Zettagram)
}

// Yottagrams returns the mass in Yg
func (m Mass) Yottagrams() float64 {
	return float64(m / Yottagram)
}

// Ronnagrams returns the mass in Rg
func (m Mass) Ronnagrams() float64 {
	return float64(m / Ronnagram)
}

// Quettagrams returns the mass in Qg
func (m Mass) Quettagrams() float64 {
	return float64(m / Quettagram)
}

// Tonnes returns the mass in t
func (m Mass) Tonnes() float64 {
	return float64(m / Tonne)
}

// Kilotonnes returns the mass in ktǂ
func (m Mass) Kilotonnes() float64 {
	return float64(m / Kilotonne)
}

// Megatonnes returns the mass in Mt
func (m Mass) Megatonnes() float64 {
	return float64(m / Megatonne)
}

// Gigatonnes returns the mass in Gt
func (m Mass) Gigatonnes() float64 {
	return float64(m / Gigatonne)
}

// Teratonnes returns the mass in Tt
func (m Mass) Teratonnes() float64 {
	return float64(m / Teratonne)
}

// Petatonnes returns the mass in Pt
func (m Mass) Petatonnes() float64 {
	return float64(m / Petatonne)
}

// Exatonnes returns the mass in Et
func (m Mass) Exatonnes() float64 {
	return float64(m / Exatonne)
}

// TroyGrains returns the mass in gr
func (m Mass) TroyGrains() float64 {
	return float64(m / TroyGrain)
}

// AvoirdupoisOunces returns the mass in oz
func (m Mass) AvoirdupoisOunces() float64 {
	return float64(m / AvoirdupoisOunce)
}

// AvoirdupoisDrams returns the mass in XXX
func (m Mass) AvoirdupoisDrams() float64 {
	return float64(m / AvoirdupoisDram)
}

// AvoirdupoisPounds returns the mass in lb
func (m Mass) AvoirdupoisPounds() float64 {
	return float64(m / AvoirdupoisPound)
}

// TroyOunces returns the mass in oz
func (m Mass) TroyOunces() float64 {
	return float64(m / TroyOunce)
}

// TroyPounds returns the mass in lb
func (m Mass) TroyPounds() float64 {
	return float64(m / TroyPound)
}

// UsStones returns the mass in st
func (m Mass) UsStones() float64 {
	return float64(m / UsStone)
}

// UkStones returns the mass in st
func (m Mass) UkStones() float64 {
	return float64(m / UkStone)
}

// UsQuarters returns the mass in qr av
func (m Mass) UsQuarters() float64 {
	return float64(m / UsQuarter)
}

// UkQuarters returns the mass in qr av
func (m Mass) UkQuarters() float64 {
	return float64(m / UkQuarter)
}

// LongHundredweights returns the mass in cwt
func (m Mass) LongHundredweights() float64 {
	return float64(m / LongHundredweight)
}

// ShortHundredweights returns the mass in cwt
func (m Mass) ShortHundredweights() float64 {
	return float64(m / ShortHundredweight)
}
