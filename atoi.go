package piscine

func Atoi(s string) int {
	o_number := 0
	c := 0
	checker := true
	a_s := []rune(s)
	pl := 1

	if s != "" {
		if byte(a_s[0]) == 45 {
			pl = -1
			a_s[0] = '0'
		} else if byte(a_s[0]) == 43 {
			a_s[0] = '0'
		}
	}
	for _, word := range a_s {
		if byte(word) >= 48 && byte(word) <= 57 {
			for i := '0'; i < word; i++ {
				c++
			}
			o_number = o_number*10 + c
			c = 0
		} else {
			checker = false
		}
	}

	if checker {
		return o_number * pl
	} else {
		return 0
	}
}
