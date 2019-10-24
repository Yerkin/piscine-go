package piscine

func Split(str, charset string) []string {
	ln := 0

	for i := range charset {
		ln = i + 1
	}

	ln2 := 0
	for i := range str {
		ln2 = i + 1
	}

	for i := 0; i < ln2-ln; i++ {
		if str[i:i+ln] == charset {
			str = str[:i] + " " + str[i+ln:]
			ln2 -= ln
		}
	}

	return SplitWhiteSpaces(str)
}
