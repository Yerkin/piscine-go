package student

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 && j == 0 {
					z01.PrintRune('o')
					if j == x-1 && i != y-1 {
						z01.PrintRune(10)
					}

				} else if i == 0 && j == x-1 {
					z01.PrintRune('o')
					if i != y-1 {
						z01.PrintRune(10)
					}

				} else if i == y-1 && j == x-1 {
					z01.PrintRune('o')

				} else if i == y-1 && j == 0 {
					z01.PrintRune('o')
					//wall
				} else if (i != 0 && i != y-1) && (j == 0) {
					z01.PrintRune('|')
					if j == x-1 {
						z01.PrintRune(10)
					}
				} else if (i != 0 && i != y-1) && (j == x-1) && (x != 1) {
					z01.PrintRune('|')
					z01.PrintRune(10)

				} else if i == 0 || i == y-1 {
					z01.PrintRune('-')
				} else {
					z01.PrintRune(32)
				}
			}
		}
		z01.PrintRune(10)
	}
}
