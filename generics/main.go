package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumsIntsOrFloats[string, int64](ints),
		SumsIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumsIntsOrFloats(ints),
		SumsIntsOrFloats(floats))

	fmt.Printf("Generic Sums, with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// Adds together the values of integers or floats m.
func SumsIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, values := range m {
		sum += values
	}

	return sum
}

// Adds together the values of Numbers m.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, values := range m {
		sum += values
	}

	return sum
}

// Adds together the values of integers m.
func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, value := range m {
		sum += value
	}

	return sum
}

// Adds together the values of floats m.
func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, value := range m {
		sum += value
	}

	return sum
}
