package piscine

func BasicAtoi2(s string) int {
	o_number := 0
	c := 0
	checker := true
	a_s := []rune(s)
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
		return o_number
	} else {
		return 0
	}
}
