package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for _, word := range str {
		z01.PrintRune(word)
	}
}
