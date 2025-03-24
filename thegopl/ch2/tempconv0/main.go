package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(((c * 9) / 5) + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius(((f - 32) * 5) / 9)
}

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	fmt.Printf("BoilingC - FreezingC = %g\n", BoilingC-FreezingC)

	fmt.Printf("FreezingC = %g, Converted to F = %g\n", FreezingC, CToF(FreezingC))

	fmt.Printf("BoilingC = %g, Converted to F = %g\n", BoilingC, CToF(BoilingC))

	fmt.Printf("CToF(BoilingC) - CToF(FreezingC) = %g\n", CToF(BoilingC)-CToF(FreezingC))

	// fmt.Printf("CToF(BoilingC) - FreezingC = %g", CToF(BoilingC) - FreezingC) // type mismatch

	c := FToC(212.0)
	fmt.Println("212°F = " + c.String())
	fmt.Printf("Also, 212°F = %v\n", c)
	fmt.Printf("Also, 212°F = %s\n", c)
	fmt.Println(c)
	fmt.Printf("Also, 212°F = %g\n", c)
	fmt.Println(float64(c))
}
