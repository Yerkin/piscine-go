package piscine

func ToUpper(s string) string {
	my_arr := []rune(s)
	for index, word := range my_arr {
		if word >= 'a' && word <= 'z' {
			my_arr[index] = word - 32
		}
	}
	return string(my_arr)
}
