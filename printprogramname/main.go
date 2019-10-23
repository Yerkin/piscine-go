package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	my_args := os.Args[0]
	my_arr := []rune(my_args)
	ln := 0
	for i := range my_args {
		ln = i
	}
	for i := 0; i <= ln; i++ {
		z01.PrintRune(rune(my_arr[i]))
	}
	z01.PrintRune('\n')
}
