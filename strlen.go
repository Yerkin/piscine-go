package piscine

func StrLen(str string) int {
	astr := []rune(str)
	a := 1
	for index := range astr {
		a = index
	}
	return a + 1
}
