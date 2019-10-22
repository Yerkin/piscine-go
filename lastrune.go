package piscine

func LastRune(s string) rune {
	my_c := 0
	for index := range s {
		my_c = index
	}
	my_arr := []rune(s)
	return my_arr[my_c]
}
