package piscine

func ToLower(s string) string {
	my_arr := []rune(s)
	for index, letter := range my_arr {
		if letter >= 'A' && letter <= 'Z' {
			my_arr[index] = letter + 32
		}
	}
	return string(my_arr)
}
