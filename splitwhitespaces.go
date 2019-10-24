package piscine

func SplitWhiteSpaces(str string) []string {
	a := 0
	ln := 0
	for _, w := range str {
		if w == '\n' || w == ' ' || w == '\t' {
			ln++
		}
	}
	ans := make([]string, ln+1)
	index := 0
	myStr := ""
	for i, w := range str {
		if w == '\n' || w == ' ' || w == '\t' {
			ans[index] = myStr
			index++
			myStr = ""
			a = i
		} else {
			myStr = myStr + string(w)
		}
	}
	ans[index] = str[a+1:]
	return ans
}
