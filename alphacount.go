package piscine

func AlphaCount(str string) int {
	my_str := []byte(str)
	my_counter := 0
	for _, letter := range my_str {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			my_counter++
		}
	}
	return my_counter
}
