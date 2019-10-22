package piscine

func NRune(s string, n int) rune {
	my_c := 0
	for index := range s {
		my_c = index
	}
	if (my_c + 1) >= n {
		my_arr := []rune(s)
		return (my_arr[n-1])
	} else {
		return 0
	}
}
