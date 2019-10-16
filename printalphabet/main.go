package main

import "github.com/01-edu/z01"

func main() {
	var i rune = 97
	var b rune = 1

	for i <=122{
		z01.PrintRune(i)
		i=i+b
	}

	z01.PrintRune('\n')

}