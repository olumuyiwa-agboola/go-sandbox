package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(((c * 9) / 5) + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius(((f - 32) * 5) / 9) }

// FToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// FToC converts a Kelvin temperature to Celsius.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
