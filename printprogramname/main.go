package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for _, w := range arg[0] {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}
