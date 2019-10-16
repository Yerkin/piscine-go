package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var a rune = 48
	var b rune = 49
	var c rune = 50
	var i rune = 1
	for a <= 55 {
		for b <= 56 {
			for c <= 57 {
				if a < b && b < c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if a < 55 {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
				c = c + i
			}
			c = 48
			b = b + i
		}
			a = a + i
			b = a + i
			c = c + i	
	}
	z01.PrintRune(10)
}
