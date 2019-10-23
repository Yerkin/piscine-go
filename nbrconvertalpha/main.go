package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	my_arr := os.Args[1:]
	ln := 0
	for i := range my_arr {
		ln = i
	}
	if ln >= 1 {
		if my_arr[0] == "--upper" {
			z01.PrintRune(' ')
			for i := 1; i <= ln; i++ {
				num := 0
				for _, w := range my_arr[i] {
					num = num*10 + int(w-'0')
				}
				if num >= 1 && num <= 26 {
					z01.PrintRune('A' + rune(num-1))
				} else {
					z01.PrintRune(' ')
				}
			}
		} else {
			for i := 0; i <= ln; i++ {
				myNum := 0
				for _, w := range my_arr[i] {
					myNum = myNum*10 + int(w-'0')
				}
				if myNum >= 1 && myNum <= 26 {
					z01.PrintRune('a' + rune(myNum-1))
				} else {
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
