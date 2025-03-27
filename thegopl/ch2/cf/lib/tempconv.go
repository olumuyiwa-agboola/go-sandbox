// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Kelvin float64
type Celsius float64
type Fahrenheit float64

const (
	// Absolute zero temperature in Celsius.
	AbsoluteZeroC Celsius = -273.15

	// Freezing temperature in Celsius.
	FreezingC Celsius = 0

	// Boiling temperature in Celsius.
	BoilingC Celsius = 0
)

// Formats a Kelvin temperature.
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

// Formats a Celsius temperature.
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// Formats a Fahrenheit temperature.
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
