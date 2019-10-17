package piscine

func StrLen(str string) int {
	a := 1
	for index := range str {
		a = index
	}
	return a + 1
}
