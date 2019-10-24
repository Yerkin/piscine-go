package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, c := range table {
		for _, j := range c {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
