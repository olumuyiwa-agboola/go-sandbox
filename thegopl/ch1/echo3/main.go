// Prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	// If we don't care about format: fmt.Println(os.Args[1:])
}
