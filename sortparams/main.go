package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	my_ar_1 := os.Args
	my_ar_2 := os.Args
	ln := 0
	for i := range my_ar_1 {
		ln = i
	}
	for i := 0; i <= ln; i++ {
		for j := 0; j <= ln; j++ {
			if my_ar_1[i] < my_ar_2[j] {
				my_ar_2[j], my_ar_1[i] = my_ar_1[i], my_ar_2[j]
			}
		}
	}
	for i := 0; i <= ln-1; i++ {
		for _, w := range my_ar_2[i] {
			z01.PrintRune(w)
		}
		z01.PrintRune(10)
	}
}
