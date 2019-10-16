package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var q int = 1
	if n < 0 {
		z01.PrintRune(45)
		q = -1
		n = n * q
	}
	if n < 10 && n >= 0 {
		z01.PrintRune(rune(n + 48))
	}
	var a int = 0
	if n > 10 {
		a = n % 10
		PrintNbr(n / 10)
		z01.PrintRune(rune(a + 48))
	}
	// z01.PrintRune(10)
}
