package piscine

func IterativeFactorial(nb int) int {
	a := 1
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 1 && nb <= 25 {
		for i := 1; i <= nb; i++ {
			a = a * i
		}
	} else {
		return 0
	}

	return a
}
