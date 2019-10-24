package main

import (
	"os"

	"github.com/01-edu/z01"
)

func chekit(x rune) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'o' || x == 'O' || x == 'u' || x == 'U' || x == 'i' || x == 'I' {
		return true
	}
	return false
}

func main() {
	myArr := os.Args[1:]
	ln := 0

	for i := range myArr {
		ln = i + 1
	}

	if ln == 0 {
		z01.PrintRune('\n')
	} else {

		myStr := myArr[0]
		for i := 1; i < ln; i++ {
			myStr = myStr + " " + myArr[i]
		}

		my_ans := []rune(myStr)

		ln2 := 0

		for i := range my_ans {
			ln2 = i
		}

		j := ln2

		for i := 0; i <= ln2; i++ {

			if j < i {
				break
			}

			if chekit(my_ans[j]) && chekit(my_ans[i]) {

			}
			j--

		}
		// fmt.Println(string(my_ans))

	}

}
