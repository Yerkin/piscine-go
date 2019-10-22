package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	my_arr := [10]int{}
	my_count := 0
	my_index := 0

	if n < 10 && n >= 0 {
		z01.PrintRune(rune(n + 48))
	} else if n >= 10 {
		for {
			if n == 0 {
				break
			}
			my_arr[my_index] = n % 10
			n /= 10
			my_count++
			my_index++
		}
		for i := 0; i <= my_count; i++ {
			for j := 0; j <= my_count; j++ {
				if my_arr[i] < my_arr[j] {
					my_arr[i], my_arr[j] = my_arr[j], my_arr[i]
				}
			}
		}
		for i := 0; i <= my_count; i++ {
			if my_arr[i] != 0 && i != 0 {
				z01.PrintRune(rune(my_arr[i] + 48))
			}
		}
	}
	// z01.PrintRune
}
