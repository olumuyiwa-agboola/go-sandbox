// Prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[0:] {
		fmt.Printf("Index: %d; Argument: %s\n", index, arg)
	}
}
