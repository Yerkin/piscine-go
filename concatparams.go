package piscine

func ConcatParams(args []string) string {
	myStr := ""
	ln := 0
	for i := range args {
		ln = i
	}
	for i := 0; i <= ln; i++ {
		if i != ln {
			myStr = myStr + args[i] + "\n"
		} else {
			myStr = myStr + args[i]
		}
	}
	return myStr
}
