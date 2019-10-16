package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for x := '0'; x <= '7'; x++ {
		for y := '1'; y <= '8'; y++ {
			for z := '2'; z <= '9'; z++ {
				if x < y && y < z {
					z01.PrintRune(x)
					z01.PrintRune(y)
					z01.PrintRune(z)
					if x < '7' {
						z01.PrintRune(44)
						z01.PrintRune(32)

					}

				}
			}
		}
	}
	z01.PrintRune(10)
}
