package eightqueens

import "github.com/01-edu/z01"

var ans [8]rune
var ban [9]bool
var ans2 [9]int

func EightQueens() {
	ok := true
	cnt := 0
	for i := 1; i <= 8; i++ {
		if ban[i] == false {
			ok = false
			//	break
		} else {
			cnt++
		}
	}

	if ok == true {
		for _, c := range ans {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')

		return
	}
	for i := '1'; i <= '8'; i++ {
		cur := 0
		for j := '1'; j <= i; j++ {
			cur++
		}
		if ban[cur] == false {
			put := true
			for j := 1; j <= cnt; j++ {
				if cur == ans2[j]-(cnt+1-j) || cur == ans2[j]+(cnt+1-j) {
					put = false
					break
				} /*
					fmt.Print(ans2[j])
					fmt.Print(" ")*/

			}
			//	fmt.Println("    <<< THIS SHIT")
			if put == true {
				ban[cur] = true
				ans[cnt] = i
				ans2[cnt+1] = cur
				EightQueens()
				ban[cur] = false
				//ans2[cnt+1] = 0

			}
		}
	}
}
