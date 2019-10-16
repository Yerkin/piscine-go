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
		if n == 0 {
			z01.PrintRune('0')
		}
		if n == 1 {
			z01.PrintRune('1')
		}
		if n == 2 {
			z01.PrintRune('2')
		}
		if n == 3 {
			z01.PrintRune('3')
		}
		if n == 4 {
			z01.PrintRune('4')
		}
		if n == 5 {
			z01.PrintRune('5')
		}
		if n == 6 {
			z01.PrintRune('6')
		}
		if n == 7 {
			z01.PrintRune('7')
		}
		if n == 8 {
			z01.PrintRune('8')
		}
		if n == 9 {
			z01.PrintRune('9')
		}
	}
	var a int = 0
	if n > 10 {
		a = n % 10
		PrintNbr(n / 10)
		if a == 0 {
			z01.PrintRune('0')
		}
		if a == 1 {
			z01.PrintRune('1')
		}
		if a == 2 {
			z01.PrintRune('2')
		}
		if a == 3 {
			z01.PrintRune('3')
		}
		if a == 4 {
			z01.PrintRune('4')
		}
		if a == 5 {
			z01.PrintRune('5')
		}
		if a == 6 {
			z01.PrintRune('6')
		}
		if a == 7 {
			z01.PrintRune('7')
		}
		if a == 8 {
			z01.PrintRune('8')
		}
		if a == 9 {
			z01.PrintRune('9')
		}
	}
	// z01.PrintRune(10)
}
