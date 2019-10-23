package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	ln := 0
	for i := range arg {
		ln = i
	}
	for i := 1; i <= ln; i++ {
		for _, w := range arg[i] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
