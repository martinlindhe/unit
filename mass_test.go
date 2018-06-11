package unit

import (
	"testing"
)

func TestMass(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptogram).Yoctograms())
	assertFloatEqual(t, 1e3, (1 * Attogram).Zeptograms())
	assertFloatEqual(t, 1e3, (1 * Femtogram).Attograms())
	assertFloatEqual(t, 1e3, (1 * Picogram).Femtograms())
	assertFloatEqual(t, 1e3, (1 * Nanogram).Picograms())
	assertFloatEqual(t, 1e3, (1 * Microgram).Nanograms())
	assertFloatEqual(t, 1e3, (1 * Milligram).Micrograms())

	assertFloatEqual(t, 1e3, (1 * Gram).Milligrams())
	assertFloatEqual(t, 1e2, (1 * Gram).Centigrams())
	assertFloatEqual(t, 1e1, (1 * Gram).Decigrams())
	assertFloatEqual(t, 1e0, (1 * Gram).Grams())
	assertFloatEqual(t, 1e-1, (1 * Gram).Decagrams())
	assertFloatEqual(t, 1e-2, (1 * Gram).Hectograms())
	assertFloatEqual(t, 1e-3, (1 * Gram).Kilograms())

	assertFloatEqual(t, 1e-3, (1 * Kilogram).Megagrams())
	assertFloatEqual(t, 1e-3, (1 * Megagram).Gigagrams())
	assertFloatEqual(t, 1e-3, (1 * Gigagram).Teragrams())
	assertFloatEqual(t, 1e-3, (1 * Teragram).Petagrams())
	assertFloatEqual(t, 1e-3, (1 * Petagram).Exagrams())
	assertFloatEqual(t, 1e-3, (1 * Exagram).Zettagrams())
	assertFloatEqual(t, 1e-3, (1 * Zettagram).Yottagrams())

	// non-SI
	assertFloatEqual(t, 1e-3, (1 * Kilogram).Tonnes())
	assertFloatEqual(t, 1e-3, (1 * Tonne).Kilotonnes())
	assertFloatEqual(t, 1e-3, (1 * Kilotonne).Megatonnes())
	assertFloatEqual(t, 1e-3, (1 * Megatonne).Gigatonnes())
	assertFloatEqual(t, 1e-3, (1 * Gigatonne).Teratonnes())
	assertFloatEqual(t, 1e-3, (1 * Teratonne).Petatonnes())
	assertFloatEqual(t, 1e-3, (1 * Petatonne).Exatonnes())

	// avoirdupois
	assertFloatEqual(t, 0.015432358352941428, (1 * Milligram).TroyGrains())
	assertFloatEqual(t, 0.0022857142857142855, (1 * TroyGrain).AvoirdupoisOunces())
	assertFloatEqual(t, 0.03657142857142857, (1 * TroyGrain).AvoirdupoisDrams())
	assertFloatEqual(t, 0.00014285714285714284, (1 * TroyGrain).AvoirdupoisPounds())

	assertFloatEqual(t, 0.07142857142857142, (1 * AvoirdupoisPound).UsStones())
	assertFloatEqual(t, 0.058775510204081644, (1 * TroyPound).UkStones())

	// https://en.wikipedia.org/wiki/Quarter_(unit)#Weight
	assertFloatEqual(t, 0.08818490487395102, (1 * Kilogram).UsQuarters())
	assertFloatEqual(t, 0.07873652220888486, (1 * Kilogram).UkQuarters())

	// https://en.wikipedia.org/wiki/Hundredweight
	assertFloatEqual(t, 50.80234544, (1 * LongHundredweight).Kilograms())
	assertFloatEqual(t, 45.359237, (1 * ShortHundredweight).Kilograms())
	assertFloatEqual(t, 1e0, (112 * AvoirdupoisPound).LongHundredweights())
	assertFloatEqual(t, 1e0, (100 * AvoirdupoisPound).ShortHundredweights())

	// troy
	assertFloatEqual(t, 0.0020833333333333333, (1 * TroyGrain).TroyOunces())
	assertFloatEqual(t, 0.0026792288807189982, (1 * Gram).TroyPounds())
}
