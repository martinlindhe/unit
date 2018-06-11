package unit

import (
	"testing"
)

func TestTemperature(t *testing.T) {

	assertFloatEqual(t, 373.15, FromCelsius(100).Kelvin())
	assertFloatEqual(t, 310.9277777777778, FromFahrenheit(100).Kelvin())
	assertFloatEqual(t, 100.0, FromKelvin(100).Kelvin())
	assertFloatEqual(t, 576.180303030303, FromNewton(100).Kelvin())
	assertFloatEqual(t, 306.4833333333333, FromDelisle(100).Kelvin())
	assertFloatEqual(t, 398.15, FromReaumur(100).Kelvin())
	assertFloatEqual(t, 449.3404761904762, FromRomer(100).Kelvin())
	assertFloatEqual(t, 55.555555555555515, FromRankine(100).Kelvin())

	assertFloatEqual(t, -173.15, FromKelvin(100).Celsius())
	assertFloatEqual(t, -279.67, FromKelvin(100).Fahrenheit())
	assertFloatEqual(t, -57.1395, FromKelvin(100).Newton())
	assertFloatEqual(t, 409.725, FromKelvin(100).Delisle())
	assertFloatEqual(t, -83.40375, FromKelvin(100).Romer())
	assertFloatEqual(t, -138.52, FromKelvin(100).Reaumur())
	assertFloatEqual(t, 180, FromKelvin(100).Rankine())
}
