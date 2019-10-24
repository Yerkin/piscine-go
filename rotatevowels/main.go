package main

import (
	"os"

	"github.com/01-edu/z01"
)

func check(x rune) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'o' || x == 'O' || x == 'u' || x == 'U' || x == 'i' || x == 'I' {
		return true
	}
	return false
}
func main() {
	arg := os.Args[1:]
	rep := []rune{}
	ans := ""
	ln := 0
	IsF := true
	for _, c := range arg {
		for _, j := range c {
			if check(j) {
				rep = append(rep, j)
				ln++
			}
		}
		if IsF {
			ans = c
			IsF = false
			continue
		}
		ans = ans + " " + c
	}
	cur := 0
	for _, c := range ans {
		if check(c) {
			z01.PrintRune(rep[ln-cur-1])
			cur++
		} else {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
