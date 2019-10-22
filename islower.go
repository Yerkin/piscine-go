package piscine

func IsLower(str string) bool {
	my_arr := []rune(str)
	ln := 0
	for i := range my_arr {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if my_arr[i] < 'a' || my_arr[i] > 'z' {
			return false
		}
	}
	return true
}
