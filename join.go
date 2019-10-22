package piscine

func Join(strs []string, sep string) string {
	my_word := ""
	ln := 0
	for i := range strs {
		ln = i
	}
	for i, w := range strs {
		if i != ln {
			my_word = my_word + w + sep
		} else {
			my_word = my_word + w
		}
	}
	return my_word
}
