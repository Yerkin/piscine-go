package main

import "github.com/01-edu/z01"

func main() {

	var i rune = 122
	var b rune = 1

	for i >= 97 {
		z01.PrintRune(i)
		i = i - b
	}

	z01.PrintRune(10)

}
