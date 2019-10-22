package piscine

func IsPrintable(str string) bool {
	my_arr := []rune(str)
	ln := 0
	for i := range my_arr {
		ln = i
	}

	for i := 0; i <= ln; i++ {
		if my_arr[i] < 32 {
			return false
		}
	}
	return true
}
