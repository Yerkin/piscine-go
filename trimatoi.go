package piscine

func TrimAtoi(s string) int {

	o_number := 0
	c := 0

	a_s := []rune(s)
	pl := 1
	minus_pos := 0
	minus_have := false
	first_num := 0

	for index, num := range a_s {
		if num == '-' {
			minus_pos = index
			minus_have = true
		}
	}

	for _, word := range a_s {
		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {
				c++
			}
			o_number = o_number*10 + c
			c = 0
		}
	}

	if minus_have {
		for index, num := range a_s {
			if num >= '0' && num <= '9' {
				first_num = index
				break
			}
		}
	}

	if first_num > minus_pos {
		pl = -1
	}

	return o_number * pl

}
