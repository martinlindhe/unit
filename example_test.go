package unit

import "fmt"

func ExampleFoot_Meters() {
	ft := 1 * Foot
	fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
	// Output: 1 feet is 0.30479999999999996 meters
}

type MyUnit int

func ExampleLength_Meters() {
	n := MyUnit(2)
	ft := Length(n) * Foot
	fmt.Println(ft.Feet(), "feet is", ft.Meters(), "meters")
	// Output: 2 feet is 0.6095999999999999 meters
}
