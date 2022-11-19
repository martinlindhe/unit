package unit

// Datasize represents a unit of data size (in bits, bit)
type Datasize Unit

// ...
const (
	// base 10 (SI prefixes)
	Bit       Datasize = 1e0
	Kilobit            = Bit * 1e3
	Megabit            = Bit * 1e6
	Gigabit            = Bit * 1e9
	Terabit            = Bit * 1e12
	Petabit            = Bit * 1e15
	Exabit             = Bit * 1e18
	Zettabit           = Bit * 1e21
	Yottabit           = Bit * 1e24
	Ronnabit           = Bit * 1e27
	Quettabit          = Bit * 1e30

	Byte       = Bit * 8
	Kilobyte   = Byte * 1e3
	Megabyte   = Byte * 1e6
	Gigabyte   = Byte * 1e9
	Terabyte   = Byte * 1e12
	Petabyte   = Byte * 1e15
	Exabyte    = Byte * 1e18
	Zettabyte  = Byte * 1e21
	Yottabyte  = Byte * 1e24
	Ronnabyte  = Byte * 1e27
	Quettabyte = Byte * 1e30

	// base 2 (IEC prefixes)
	Kibibit = Bit * 1024
	Mebibit = Kibibit * 1024
	Gibibit = Mebibit * 1024
	Tebibit = Gibibit * 1024
	Pebibit = Tebibit * 1024
	Exbibit = Pebibit * 1024
	Zebibit = Exbibit * 1024
	Yobibit = Zebibit * 1024

	Kibibyte = Byte * 1024
	Mebibyte = Kibibyte * 1024
	Gibibyte = Mebibyte * 1024
	Tebibyte = Gibibyte * 1024
	Pebibyte = Tebibyte * 1024
	Exbibyte = Pebibyte * 1024
	Zebibyte = Exbibyte * 1024
	Yobibyte = Zebibyte * 1024
)

// Bits returns the datasize in bit
func (b Datasize) Bits() float64 {
	return float64(b)
}

// Kilobits returns the datasize in kbit
func (b Datasize) Kilobits() float64 {
	return float64(b / Kilobit)
}

// Megabits returns the datasize in Mbit
func (b Datasize) Megabits() float64 {
	return float64(b / Megabit)
}

// Gigabits returns the datasize in Gbit
func (b Datasize) Gigabits() float64 {
	return float64(b / Gigabit)
}

// Terabits returns the datasize in Tbit
func (b Datasize) Terabits() float64 {
	return float64(b / Terabit)
}

// Petabits returns the datasize in Pbit
func (b Datasize) Petabits() float64 {
	return float64(b / Petabit)
}

// Exabits returns the datasize in Ebit
func (b Datasize) Exabits() float64 {
	return float64(b / Exabit)
}

// Zettabits returns the datasize in Zbit
func (b Datasize) Zettabits() float64 {
	return float64(b / Zettabit)
}

// Yottabits returns the datasize in Ybit
func (b Datasize) Yottabits() float64 {
	return float64(b / Yottabit)
}

// Ronnabits returns the datasize in Rbit
func (b Datasize) Ronnabits() float64 {
	return float64(b / Ronnabit)
}

// Quettabits returns the datasize in Qbit
func (b Datasize) Quettabits() float64 {
	return float64(b / Quettabit)
}

// Bytes returns the datasize in B
func (b Datasize) Bytes() float64 {
	return float64(b / Byte)
}

// Kilobytes returns the datasize in kB
func (b Datasize) Kilobytes() float64 {
	return float64(b / Kilobyte)
}

// Megabytes returns the datasize in MB
func (b Datasize) Megabytes() float64 {
	return float64(b / Megabyte)
}

// Gigabytes returns the datasize in GB
func (b Datasize) Gigabytes() float64 {
	return float64(b / Gigabyte)
}

// Terabytes returns the datasize in TB
func (b Datasize) Terabytes() float64 {
	return float64(b / Terabyte)
}

// Petabytes returns the datasize in PB
func (b Datasize) Petabytes() float64 {
	return float64(b / Petabyte)
}

// Exabytes returns the datasize in EB
func (b Datasize) Exabytes() float64 {
	return float64(b / Exabyte)
}

// Zettabytes returns the datasize in ZB
func (b Datasize) Zettabytes() float64 {
	return float64(b / Zettabyte)
}

// Yottabytes returns the datasize in YB
func (b Datasize) Yottabytes() float64 {
	return float64(b / Yottabyte)
}

// Ronnabytes returns the datasize in RB
func (b Datasize) Ronnabytes() float64 {
	return float64(b / Ronnabyte)
}

// Quettabytes returns the datasize in RB
func (b Datasize) Quettabytes() float64 {
	return float64(b / Quettabyte)
}

// Kibibits returns the datasize in Kibit
func (b Datasize) Kibibits() float64 {
	return float64(b / Kibibit)
}

// Mebibits returns the datasize in Mibit
func (b Datasize) Mebibits() float64 {
	return float64(b / Mebibit)
}

// Gibibits returns the datasize in Gibit
func (b Datasize) Gibibits() float64 {
	return float64(b / Gibibit)
}

// Tebibits returns the datasize in Tibit
func (b Datasize) Tebibits() float64 {
	return float64(b / Tebibit)
}

// Pebibits returns the datasize in Pibit
func (b Datasize) Pebibits() float64 {
	return float64(b / Pebibit)
}

// Exbibits returns the datasize in Eibit
func (b Datasize) Exbibits() float64 {
	return float64(b / Exbibit)
}

// Zebibits returns the datasize in Zibit
func (b Datasize) Zebibits() float64 {
	return float64(b / Zebibit)
}

// Yobibits returns the datasize in Yibit
func (b Datasize) Yobibits() float64 {
	return float64(b / Yobibit)
}

// Kibibytes returns the datasize in KiB
func (b Datasize) Kibibytes() float64 {
	return float64(b / Kibibyte)
}

// Mebibytes returns the datasize in MiB
func (b Datasize) Mebibytes() float64 {
	return float64(b / Mebibyte)
}

// Gibibytes returns the datasize in GiB
func (b Datasize) Gibibytes() float64 {
	return float64(b / Gibibyte)
}

// Tebibytes returns the datasize in TiB
func (b Datasize) Tebibytes() float64 {
	return float64(b / Tebibyte)
}

// Pebibytes returns the datasize in PiB
func (b Datasize) Pebibytes() float64 {
	return float64(b / Pebibyte)
}

// Exbibytes returns the datasize in EiB
func (b Datasize) Exbibytes() float64 {
	return float64(b / Exbibyte)
}

// Zebibytes returns the datasize in ZiB
func (b Datasize) Zebibytes() float64 {
	return float64(b / Zebibyte)
}

// Yobibytes returns the datasize in YiB
func (b Datasize) Yobibytes() float64 {
	return float64(b / Yobibyte)
}
