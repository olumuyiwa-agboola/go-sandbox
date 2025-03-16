// Prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func main() {
	res1 := testing.Benchmark(BenchmarkFirstVersion)
	fmt.Printf("Memory allocations : %d \n", res1.MemAllocs)
	fmt.Printf("Number of bytes allocated: %d \n", res1.Bytes)
	fmt.Printf("Number of run: %d \n", res1.N)
	fmt.Printf("Time taken: %s \n", res1.T)

	fmt.Println("----------------------------------------------------------------")

	res2 := testing.Benchmark(BenchmarkSecondVersion)
	fmt.Printf("Memory allocations : %d \n", res2.MemAllocs)
	fmt.Printf("Number of bytes allocated: %d \n", res2.Bytes)
	fmt.Printf("Number of run: %d \n", res2.N)
	fmt.Printf("Time taken: %s \n", res2.T)
}

func FirstVersion() string {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	return s
}

var result string

func BenchmarkFirstVersion(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = FirstVersion()
	}

	result = s
}

func SecondVersion() string {
	return strings.Join(os.Args[1:], " ")
}

func BenchmarkSecondVersion(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = SecondVersion()
	}

	result = s
}
