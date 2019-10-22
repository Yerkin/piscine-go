package piscine

func find_len(my_s string) int {
	my_len := 0
	for index := range my_s {
		my_len = index
	}
	return my_len + 1
}

func Index(s string, toFind string) int {

	len_a := find_len(s)
	len_b := find_len(toFind)

	for i := 0; i < len_a; i++ {
		if len_b != 0 && s[i] == toFind[0] {
			ok := true
			for j := 1; j < len_b; j++ {
				if s[i+j] != toFind[j] {
					ok = false
					break
				}
			}
			if ok == true {
				return i
			}
		}

	}
	return -1
}
