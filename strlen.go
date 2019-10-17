package piscine

func StrLen(str string) int {
	a := 1
	for index, _ := range str {
		a = index
	}
	return a+1
}
